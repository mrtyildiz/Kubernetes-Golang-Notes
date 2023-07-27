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

	// Secret objesini oluştur
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: "default",
		},
		Data: map[string][]byte{
			"username": []byte("my-username"),
			"password": []byte("my-password"),
		},
		Type: corev1.SecretTypeOpaque,
	}

	// Secret objesini Kubernetes kümesine gönder
	secretsClient := clientset.CoreV1().Secrets(secret.Namespace)
	result, err := secretsClient.Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Secret %q başarıyla oluşturuldu.\n", result.GetObjectMeta().GetName())
}
