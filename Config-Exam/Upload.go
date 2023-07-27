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

	// ConfigMap objesini al
	configMapName := "my-configmap"
	configMapNamespace := "default"
	configMap, err := clientset.CoreV1().ConfigMaps(configMapNamespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// ConfigMap verilerini güncelle
	configMap.Data["key1"] = "new-value1"
	configMap.Data["key3"] = "value3"

	// ConfigMap'ı güncelle
	result, err := clientset.CoreV1().ConfigMaps(configMap.Namespace).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ConfigMap %q başarıyla güncellendi.\n", result.GetObjectMeta().GetName())
}
