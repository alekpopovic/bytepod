package main

import (
	"net/http"
	"time"

	"k8s.io/client-go/kubernetes"
)

type ExternalAPIPayload struct {
	EventType string    `json:"event_type"`
	PodName   string    `json:"pod_name"`
	Namespace string    `json:"namespace"`
	Timestamp time.Time `json:"timestamp"`
	PodPhase  string    `json:"pod_phase"`
	NodeName  string    `json:"node_name"`
}

type KubernetesPlugin struct {
	clientset   *kubernetes.Clientset
	httpClient  *http.Client
	externalURL string
}
