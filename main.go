package main

import (
	"context"
	"log"
	"os"
)

func main() {
	externalAPIURL := os.Getenv("EXTERNAL_API_URL")
	if externalAPIURL == "" {
		log.Fatal("EXTERNAL_API_URL environment variable is required")
	}

	namespace := os.Getenv("WATCH_NAMESPACE")

	plugin, err := NewKubernetesPlugin(externalAPIURL)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes plugin: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("Starting Kubernetes plugin, external API: %s", externalAPIURL)

	if err := plugin.WatchPods(ctx, namespace); err != nil {
		log.Fatalf("Error watching pods: %v", err)
	}
}
