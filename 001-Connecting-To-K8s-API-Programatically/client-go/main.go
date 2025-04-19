package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string

	// Attempt to locate the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback if home directory not found
		log.Printf("Warning: Could not find home directory: %v. --kubeconfig flag will be required.", err)
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file (required if home dir not found)")
	} else {
		// Set default path to ~/.kube/config
		defaultPath := filepath.Join(home, ".kube", "config")
		kubeconfig = flag.String("kubeconfig", defaultPath, "(optional) absolute path to the kubeconfig file")
	}

	// Parse command-line flags
	flag.Parse()

	// Load configuration from kubeconfig path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a new Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %v", err)
	}

	// Use clientset to list Pods in the default namespace
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	// Output the names of the Pods found
	fmt.Println("Pods in default namespace:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}
