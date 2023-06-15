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

// returns clientSet to interact with kubernetes using kubeconfig
func getClientSet() *kubernetes.Clientset {
	var kubeconfig *string

	// reads kubeconfig from the home dir/ absolute path
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig location")
	} else {
		kubeconfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")

	}
	flag.Parse()

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

// Creates a kubernetes deployment resource
func createWebDeployment(namespace string, replicasCount int32, clientSet *kubernetes.Clientset, deploymentName string, key string, value string) {
	// creating deployment struct
	deployment := &appsv1.Deployment{

		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicasCount,
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
	//Call AppsV1() method to return appsv1.AppsV1Interface and Deployment method for the DeploymentInterface
	deployment, err := clientSet.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created deployment %s.\n", deployment.GetName())
}

// check if the replicas has all pods ready
func checkForReadyReplicas(replicasCount int32, deploymentName string, clientSet *kubernetes.Clientset, namespace string) {
	//waiting to the replicas to get ready.
	for {
		deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		currentReadyReplicas := (*deployment).Status.ReadyReplicas
		fmt.Println("Current Status ReadyReplicas = ", currentReadyReplicas)
		if currentReadyReplicas == replicasCount {
			break
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("All Pods are ready")
}

// creating a kubernetes service resource
func createService(clientSet *kubernetes.Clientset, namespace string, serviceName string, key string, value string, servicePort int32) {
	//defining service struct
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
	//Call CoreV1() method to return CoreV1Interface and Services method for the ServiceInterface
	service, err := clientSet.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Service Created: %q\n", service.ObjectMeta.Name)
}

func createIngress(clientSet *kubernetes.Clientset, namespace string, serviceName string, path string, host string) {
	prefix := networkingv1.PathTypePrefix
	//Define ingress struct
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myingress",
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: host,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									PathType: &prefix,
									Path:     path,
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: serviceName,
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

	//variables
	namespace := "default"

	//labels and selector
	key := "appname"
	value := "myapp"

	deploymentName := "mydeployment"
	replicasCount := int32(2)
	//Creating a deployment
	createWebDeployment(namespace, replicasCount, clientSet, deploymentName, key, value)

	//Wait for replicas to come online
	checkForReadyReplicas(replicasCount, deploymentName, clientSet, namespace)

	serviceName := "myservice"
	servicePort := int32(80)
	//create service
	createService(clientSet, namespace, serviceName, key, value, servicePort)

	path := "/home"
	host := "example.something.com"
	//creating ingress resource
	createIngress(clientSet, namespace, serviceName, path, host)
}
