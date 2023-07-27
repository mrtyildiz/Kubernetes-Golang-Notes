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
	kubeconfig := `C:\Users\Procenne-Murat\.kube\config`

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

	// İstemciyi kullanarak Kubernetes kümesine istek yapabilirsiniz
	// Örnek olarak tüm Pod'ları listeleme:
	pods, err := clientset.CoreV1().Pods("new-namespace").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
