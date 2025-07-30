package main

import (
	"context"
	"fmt"

	"github.com/justmamadou/go-kubernetes/k8s"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		err       error
		//labels    map[string]string
		//deployments []appsv1.Deployment
	)
	ctx := context.Background()

	clientSet, err = k8s.GetClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}
	//fmt.Printf("%+v\n", clientSet)
	/*
			labels, err = k8s.Deploy(ctx, clientSet)
			if err != nil {
				fmt.Printf("Error deploying: %v\n", err)
				return
			}
			fmt.Printf("Deployment labels: %v\n", labels)
			fmt.Println("Deployment successful!")


		err = k8s.DeleteDeployment(ctx, clientSet, "my-app", "default")
		if err != nil {
			fmt.Printf("Error deleting deployment: %v\n", err)
			return
		}
		fmt.Println("Deployment deleted successfully!")

		/*
			deployments, err = k8s.ListDeployments(ctx, clientSet, "uat")
			if err != nil {
				fmt.Printf("Error listing deployments: %v\n", err)
				return
			}

			for _, dep := range deployments {
				fmt.Printf("Deployment Name: %s\n", dep.Name)
			}
	*/

	deploy, err := k8s.GetDeployment(ctx, clientSet, "aggregatorcachin", "uat")
	if err != nil {
		fmt.Printf("Error getting deployment: %v\n", err)
		return
	}

	fmt.Printf("Got deployment: %s\nAvailable replicas: %d\n", deploy.Name, deploy.Status.AvailableReplicas)

}
