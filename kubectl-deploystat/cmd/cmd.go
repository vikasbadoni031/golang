package cmd

import (
	"context"
	"fmt"
	"kubectl-deploystat/client"
	"strings"

	cobra "github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type options struct {
	namespace string
}

func Deploystat() *cobra.Command {
	opt := options{}
	cobraVar := &cobra.Command{
		Use:     "deploystat",
		Aliases: []string{"deploystat"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDeployStat(args[0], opt)
		},
		Short:   "Get deployment stats",
		Long:    "Get all deployments-> Container stats for CPU and Memory",
		Version: "1.0",
		Example: "kubectl deploystat -n default",
	}
	cobraVar.Flags().StringVarP(&opt.namespace, "namespace", "n", "default", "namespace to check in")
	return cobraVar
}

func getDeployStat(inputDeploy string, opt options) error {

	clientset := client.ClientSet()
	deployments, err := clientset.AppsV1().Deployments(opt.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, deployment := range deployments.Items {
		deployName := deployment.ObjectMeta.Name
		if !strings.Contains(deployName, inputDeploy) {
			continue
		}
		if err != nil {
			return err
		}
		fmt.Print(deployName, "   ")
		for _, container := range deployment.Spec.Template.Spec.Containers {
			if len(container.Resources.Requests) == 0 {
				fmt.Print("No resources defined for this deployment")
			}
			for resourceName, quantity := range container.Resources.Requests {
				fmt.Print(string(resourceName), "=", quantity.ToDec(), "   ") //method to convert

			}
			fmt.Print("\n")
		}

	}
	return nil
}
