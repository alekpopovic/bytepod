package main

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

func (kp *KubernetesPlugin) WatchPods(ctx context.Context, namespace string) error {
	watchlist := fields.Everything()
	if namespace == "" {
		namespace = metav1.NamespaceAll
	}

	watcher, err := kp.clientset.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{
		FieldSelector: watchlist.String(),
	})
	if err != nil {
		return fmt.Errorf("failed to create pod watcher: %w", err)
	}
	defer watcher.Stop()

	log.Printf("Started watching pods in namespace: %s", namespace)

	for {
		select {
		case event, ok := <-watcher.ResultChan():
			if !ok {
				log.Println("Pod watcher channel closed, restarting...")
				return kp.WatchPods(ctx, namespace)
			}

			if pod, ok := event.Object.(*v1.Pod); ok {
				log.Printf("Pod event: %s - %s/%s", event.Type, pod.Namespace, pod.Name)
				kp.handlePodEvent(event.Type, pod)
			}

		case <-ctx.Done():
			log.Println("Context cancelled, stopping pod watcher")
			return ctx.Err()
		}
	}
}
