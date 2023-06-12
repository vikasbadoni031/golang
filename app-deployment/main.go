package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
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

	// Now we are passing the value of kubeconfig. by using *
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	namespace := "default"
	replicasCount := int32(2)
	deployment := &appsv1.Deployment{

		ObjectMeta: metav1.ObjectMeta{
			Name: "mydeployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicasCount, //not we are using a function here to convert int32 to an address
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"appname": "myapp",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"appname": "myapp",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "container1",
							Image: "nginx:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "webcontainer",
									Protocol:      apiv1.ProtocolTCP,
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
	//fmt.Println("zindagi")
	//fmt.Printf("%p", *deployment)
	// as the second arg to Create func is "*v1.Deployment" anything which is passed to it should be an address.
	// can be seen using "fmt.Printf("%p", deployment)" which will print an address
	result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	//retured value result is an address too, look at the signature of the create func in deployment interface
	if err != nil {
		panic(err)
	}
	fmt.Printf("%p", result)

	//func (*Deployment) Size , size method takes pointer to deployment type as a receiver argument
	//https://pkg.go.dev/k8s.io/api/apps/v1#Deployment.Size
	fmt.Println(result.Size())

	// here GetName is an embbeded method from ObjectMeta which is a embbeded struct in the deployment Type.
	// struct deployment{ objectMeta--> func GetName() }
	fmt.Printf("Created deployment %s.\n", result.GetName())

}

func int32Ptr(i int32) *int32 { return &i }
