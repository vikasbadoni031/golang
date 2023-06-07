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

	deployment := &appsv1.Deployment{

		ObjectMeta: metav1.ObjectMeta{
			Name: "mydeployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
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

	// as the second arg to Create func is "*v1.Deployment" anything which is passed to it should be an address.
	result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("1111")

	fmt.Println(result.Size())

	fmt.Println("2222")

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

}

func int32Ptr(i int32) *int32 { return &i }
