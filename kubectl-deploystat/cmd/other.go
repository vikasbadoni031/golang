package cmd

// Other options
// func getPodStat() error {
// 	clientset := client.ClientSet()
// 	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, pod := range pods.Items {
// 		//podname
// 		fmt.Print(pod.ObjectMeta.Name, "   ")
// 		podDetails, err := clientset.CoreV1().Pods("default").Get(context.TODO(), pod.ObjectMeta.Name, metav1.GetOptions{})
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, container := range podDetails.Spec.Containers {
// 			for resourceName, quantity := range container.Resources.Requests {
// 				fmt.Print(string(resourceName), "=", quantity.ToDec(), "   ") //method to convert

// 			}
// 			fmt.Print("\n")
// 		}
// 	}
// 	return nil
// }
