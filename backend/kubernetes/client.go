package kubernetes

import (
	"os"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func NewDynamicClient() (dynamic.Interface, error) {

	// Uses ~/.kube/config by default
	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return dynamic.NewForConfig(config)
}