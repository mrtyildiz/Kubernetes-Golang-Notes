package main

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// Kubernetes yapılandırma dosyasının yolu
	kubeconfig := "C:\\Users\\Procenne-Murat\\.kube\\config"

	// Kubernetes istemci oluşturma
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Kubernetes istemcisini oluştur
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Pod'ı sil
	podName := "nginx-deployment-7dbf6d9ddc-jwzn2"
	podNamespace := "default"
	err = clientset.CoreV1().Pods(podNamespace).Delete(context.TODO(), podName, v1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pod %q başarıyla silindi.\n", podName)

	// Deployment'ı sil
	deploymentName := "nginx-deployment"
	deploymentNamespace := "default"
	err = clientset.AppsV1().Deployments(deploymentNamespace).Delete(context.TODO(), deploymentName, v1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %q başarıyla silindi.\n", deploymentName)
}
