# bytepod

## Plugin Architecture: 

- This Kubernetes plugin uses the client-go library to watch Pod events and forwards them to an external API endpoint

## Configuration:

- Requires EXTERNAL_API_URL environment variable for the external API endpoint

- Optional WATCH_NAMESPACE environment variable to limit watching to a specific namespace

- Automatically detects in-cluster config or falls back to local kubeconfig

## Key Features:

- Watches for Pod creation, updates, and deletion events

- Sends structured JSON payloads to external API with pod metadata

- Includes proper error handling and retry logic

- Uses context for graceful shutdown

- Supports both in-cluster and local development environments

## Deployment Requirements:

- Kubernetes RBAC permissions to watch pods

- Network access to external API endpoint

- Can be deployed as a Deployment or Job in the cluster


## Dependencies (go.mod):

- require k8s.io/client-go

- require k8s.io/api

- require k8s.io/apimachinery


# Usage: 

- Build as container image and deploy to Kubernetes cluster with appropriate ServiceAccount and RBAC permissions