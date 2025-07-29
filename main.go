package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		err       error
		labels    map[string]string
	)

	clientSet, err = getClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}
	//fmt.Printf("%+v\n", clientSet)
	labels, err = deploy(context.Background(), clientSet)
	if err != nil {
		fmt.Printf("Error deploying: %v\n", err)
		return
	}
	fmt.Printf("Deployment labels: %v\n", labels)
	fmt.Println("Deployment successful!")
}

func getClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	} // This function is intended to return a client instance.
	// You can implement the logic to create and return a client here.
	return clientset, nil
}

func deploy(ctx context.Context, clientset *kubernetes.Clientset) (map[string]string, error) {
	var deployment *appsv1.Deployment
	fileName, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		return nil, err
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(fileName), nil, nil)
	if err != nil {
		return nil, err
	}

	deployment, ok := obj.(*appsv1.Deployment)
	if !ok {
		return nil, fmt.Errorf("expected Deployment object, got %T", obj)
	}

	apps, err := clientset.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return apps.Spec.Template.Labels, nil
}
