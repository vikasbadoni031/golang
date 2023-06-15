package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getClientSet() *kubernetes.Clientset {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		// The return value is the address of a string variable that stores the value of the flag. The reason the String function is returning the address in bcoz string type is pass by value in golang and if the function doesnt returns value instread of address we will end up modifying a different var in the memory.
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig location")
	} else {
		kubeconfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")

	}
	//fmt.Println(kubeconfig)
	//fmt.Println(*kubeconfig)
	flag.Parse()

	// Now we are passing the value of kubeconfig. by using *, bcoz the above returned kubeconfig is an address
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientSet
}

func createWebDeployment(namespace string, replicasCount int32, clientSet *kubernetes.Clientset, deploymentName string, key string, value string) {
	deployment := &appsv1.Deployment{ //& bcoz u need to pass address in the next call.

		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicasCount, //convert int32 to an address as thats what a function expects
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					key: value,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						key: value,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container1",
							Image: "nginx:latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "webcontainer",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	//fmt.Printf("%p", deployment)
	//fmt.Printf("%p", *deployment)
	// as the second arg to Create func is "*v1.Deployment" anything which is passed to it should be an address. henec we write this &appsv1.Deployment
	// can be seen using "fmt.Printf("%p", deployment)" which will print an address
	deployment, err := clientSet.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	//retured value result is an address too, look at the signature of the create func in deployment interface
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%p", result)

	//func (*Deployment) Size , size method takes pointer to deployment type as a receiver argument
	//https://pkg.go.dev/k8s.io/api/apps/v1#Deployment.Size
	//fmt.Println(result.Size())

	// here GetName is an embbeded method from ObjectMeta which is a embbeded struct in the deployment Type.
	// struct deployment{ objectMeta--> func GetName() }
	fmt.Printf("Created deployment %s.\n", deployment.GetName())
	//return deployment
}

func checkForReadyReplicas(replicasCount int32, deploymentName string, clientSet *kubernetes.Clientset, namespace string) {

	//status field of deployment struct, and readyReplicas field of the DeploymentStatus nested struct
	//currentReadyReplicas := (*deployment).Status.ReadyReplicas //calling a method on pointer should be done inside a bracket else it will consider the whole string to be a pointer
	//currentReadyReplicas := deployment.Status.ReadyReplicas // we can also do this.
	//fmt.Println(deployment.Status.ReadyReplicas)

	//waiting to the replicas to get ready.
	for {
		// we need to recrete this client for each iteration in loop else we are not able to get the current status of the replcaset.
		deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		currentReadyReplicas := (*deployment).Status.ReadyReplicas
		fmt.Println("Current status ReadyReplicas", currentReadyReplicas)
		if currentReadyReplicas == replicasCount {
			break
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Replicas are ready")
}

func createService(clientSet *kubernetes.Clientset, namespace string, serviceName string, key string, value string, servicePort int32) {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				key: value,
			},
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Port:     servicePort,
					Protocol: corev1.ProtocolTCP,
				},
			},
		},
	}
	service, err := clientSet.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Service Created", service.ObjectMeta.Name) //another option to get name
	fmt.Printf("Service Created: %q\n", service.ObjectMeta.Name)
}

func createIngress(clientSet *kubernetes.Clientset, namespace string, serviceName string) {
	prefix := networkingv1.PathTypePrefix
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myingress",
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "example.something.com",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{ //expected value is an address hence we added '&'
							Paths: []networkingv1.HTTPIngressPath{ //wheneven there is a list, u need to add additional brackets
								{
									PathType: &prefix,
									Path:     "/home",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "myservice",
											Port: networkingv1.ServiceBackendPort{
												Name: "http",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	ingress, err := clientSet.NetworkingV1().Ingresses(namespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ingress Created: %q\n", ingress.GetName())
}

func main() {
	clientSet := getClientSet() //get authentication info via kubeconfig for the cluster

	namespace := "default"
	replicasCount := int32(2)
	//labels and selector
	key := "appname"
	value := "myapp"

	deploymentName := "mydeployment"
	serviceName := "myservice"
	servicePort := int32(80)

	createWebDeployment(namespace, replicasCount, clientSet, deploymentName, key, value)
	//fmt.Printf("%p", deployment) // it is address
	checkForReadyReplicas(replicasCount, deploymentName, clientSet, namespace)

	createService(clientSet, namespace, serviceName, key, value, servicePort)

	//creating ingress resource
	createIngress(clientSet, namespace, serviceName)
}
