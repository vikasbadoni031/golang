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

//k8s.io/client-go/kubernetes/typed/core/v1
/*

 */

func main() {
	var kubeconfig, namespace *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	namespace = flag.String("namespace", "cb-kcdev01-apps", "Namespace to query for")
	flag.Parse()

	// use the current context in kubeconfig
	//fmt.Println(kubeconfig)
	//config, err := clientcmd.BuildConfigFromFlags("", "/Users/vbadoni/.kube/config")

	// BuildConfigFromFlags(masterUrl, kubeconfigPath string) (*restclient.Config, error)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Namespaces().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, p := range pods.Items {
		if p.Status.Phase == "Running" {
			fmt.Printf("%s - %s\n", p.ObjectMeta.Name, p.ObjectMeta.Annotations["kubernetes.io/psp"])
		}
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
