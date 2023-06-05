package main

import (
	"flag"
	"fmt"
	"path/filepath"

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
	fmt.Println(kubeconfig)
	fmt.Println(*kubeconfig)

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

}
https://github.com/PacktPublishing/Go-for-DevOps/blob/rev0/chapter/14/workloads/main.go
https://github.com/kubernetes/client-go/blob/v0.27.2/examples/create-update-delete-deployment/main.go