package k8s

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetClient() (*kubernetes.Clientset, error) {
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

func Deploy(ctx context.Context, clientset *kubernetes.Clientset) (map[string]string, error) {
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

func DeleteDeployment(ctx context.Context, clientset *kubernetes.Clientset, name, namespace string) error {
	deployment, err := GetDeployment(ctx, clientset, name, namespace)
	if err != nil {
		return err
	}
	return clientset.AppsV1().Deployments(namespace).Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}

func ListDeployments(ctx context.Context, clientset *kubernetes.Clientset, namespace string) ([]appsv1.Deployment, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments.Items, nil
}

func GetDeployment(ctx context.Context, clientset *kubernetes.Clientset, name, namespace string) (*appsv1.Deployment, error) {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}
