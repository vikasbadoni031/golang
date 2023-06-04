package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig, namespace *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	namespace = flag.String("namespace", "default", "Namespace to query for")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error)
	}

	// creating clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//Working with Configmaps

	//Listing CM and getting data
	configMaps, err := clientset.CoreV1().ConfigMaps(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, configMap := range configMaps.Items {
		fmt.Println(configMap.Data)            //data
		fmt.Println(configMap.ObjectMeta.Name) //Name

	}

	//get CM
	configMapGet, err := clientset.CoreV1().ConfigMaps(*namespace).Get(context.TODO(), "somthing", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(configMapGet.ObjectMeta.Annotations)

	//get pods for a particular release based on label release and have cpu request more that 200m.

	podList, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range podList.Items {
		val := pod.ObjectMeta.Labels["app.kubernetes.io/release"]
		if val == "service" {
			for _, container := range pod.Spec.Containers {
				cpuValueMV := container.Resources.Requests.Cpu().MilliValue()
				if cpuValueMV > 2 {
					fmt.Println(pod.ObjectMeta.Name, cpuValueMV)
				}
			}
		}
	}

	// List all deployments Names and no of replicas
	deployments, _ := clientset.AppsV1().Deployments(*namespace).List(context.TODO(), metav1.ListOptions{})
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.ObjectMeta.Name, "--Replicas->", *deployment.Spec.Replicas)
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
