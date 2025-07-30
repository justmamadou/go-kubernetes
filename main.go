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
	*/

	err = k8s.DeleteDeployment(ctx, clientSet, "my-app")
	if err != nil {
		fmt.Printf("Error deleting deployment: %v\n", err)
		return
	}
	fmt.Println("Deployment deleted successfully!")
}
