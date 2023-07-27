package main

import (
	"context"
	"fmt"
	"io/ioutil"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

	// app.yaml dosyasını oku
	yamlFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		panic(err.Error())
	}

	// YAML dosyasındaki verileri Deployment kaynağına dönüştür
	var deployment appsv1.Deployment
	if err := yaml.Unmarshal(yamlFile, &deployment); err != nil {
		panic(err.Error())
	}

	// Deployment kaynağını Kubernetes kümesine gönder
	deploymentsClient := clientset.AppsV1().Deployments(deployment.Namespace)
	result, err := deploymentsClient.Create(context.TODO(), &deployment, v1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %q başarıyla oluşturuldu.\n", result.GetObjectMeta().GetName())
}
