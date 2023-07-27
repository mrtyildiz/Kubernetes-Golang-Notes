package main

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	corev1 "k8s.io/api/core/v1"
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

	// ConfigMap objesini oluştur
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-configmap",
			Namespace: "default",
		},
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// ConfigMap objesini Kubernetes kümesine gönder
	configMapsClient := clientset.CoreV1().ConfigMaps(configMap.Namespace)
	result, err := configMapsClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ConfigMap %q başarıyla oluşturuldu.\n", result.GetObjectMeta().GetName())
}
