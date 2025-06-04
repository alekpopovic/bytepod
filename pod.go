package main

import (
	"log"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (kp *KubernetesPlugin) handlePodEvent(eventType watch.EventType, pod *v1.Pod) {
	payload := ExternalAPIPayload{
		EventType: string(eventType),
		PodName:   pod.Name,
		Namespace: pod.Namespace,
		Timestamp: time.Now(),
		PodPhase:  string(pod.Status.Phase),
		NodeName:  pod.Spec.NodeName,
	}

	if err := kp.sendToExternalAPI(payload); err != nil {
		log.Printf("Error sending to external API: %v", err)
	}
}
