package main

import (
	"context"
	"fmt"
	"log"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//loading config file
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client config: %v", err)
	}
	// creating clientset to communicate with the api server
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %v", err)
	}
	// defining a pod object
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-pod",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx:latest",
				},
			},
		},
	}
	createdPod, err := createPod(clientset, "default", pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error while creating pod: %v", err)
	}
	fmt.Printf("Pod created successfully\n Name: %v \n ", createdPod.Name)
}
func createPod(clientset *kubernetes.Clientset, namespace string, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error) {
	createdPod, err := clientset.CoreV1().Pods(namespace).Create(context.Background(), pod, opts)
	if err != nil {
		return nil, err
	}
	return createdPod, nil
}
