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

	// ConfigMap'ı sil
	configMapName := "my-configmap"
	configMapNamespace := "default"
	err = clientset.CoreV1().ConfigMaps(configMapNamespace).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ConfigMap %q başarıyla silindi.\n", configMapName)
}
