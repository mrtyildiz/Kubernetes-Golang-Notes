package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	corev1 "k8s.io/api/core/v1"
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

	// Etkinlik izleyicisi oluştur
	eventListWatcher := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"events",
		"default",
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				event := obj.(*corev1.Event)
				fmt.Printf("Yeni bir etkinlik eklendi: %s: %s\n", event.InvolvedObject.Name, event.Message)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				newEvent := newObj.(*corev1.Event)
				fmt.Printf("Etkinlik güncellendi: %s: %s\n", newEvent.InvolvedObject.Name, newEvent.Message)
			},
			DeleteFunc: func(obj interface{}) {
				event := obj.(*corev1.Event)
				fmt.Printf("Bir etkinlik silindi: %s: %s\n", event.InvolvedObject.Name, event.Message)
			},
		},
	)

	stopChannel := make(chan struct{})
	defer close(stopChannel)

	// Etkinlik izleyiciyi başlat
	eventInformer := cache.NewSharedIndexInformer(eventListWatcher, &corev1.Event{}, 0, cache.Indexers{})
	go eventInformer.Run(stopChannel)

	// Uygulamayı belirli bir süre boyunca çalıştır
	time.Sleep(5 * time.Minute)
}
