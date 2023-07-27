package main

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// Secret'ı sil
	secretName := "my-secret"
	secretNamespace := "default"
	err = clientset.CoreV1().Secrets(secretNamespace).Delete(context.TODO(), secretName, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Secret %q başarıyla silindi.\n", secretName)
}
