package main

import (
	"fmt"
	"net/http"
	"time"

	"k8s.io/client-go/kubernetes"
)

func NewKubernetesPlugin(externalURL string) (*KubernetesPlugin, error) {
	config, err := getKubernetesConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get kubernetes config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes clientset: %w", err)
	}

	return &KubernetesPlugin{
		clientset:   clientset,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		externalURL: externalURL,
	}, nil
}
