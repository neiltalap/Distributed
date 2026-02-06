# Ingress

An **Ingress** exposes HTTP and HTTPS routes from outside the cluster to services within the cluster.

## Manifests

| File | Description |
|------|-------------|
| `ingress-basic.yaml` | Standard Ingress resource for path-based routing. |

## Prerequisites

- An **Ingress Controller** must be running in your cluster (e.g., NGINX Ingress Controller).
    - ***Minikube***: `minikube addons enable ingress`

## Usage

1. Apply the Ingress manifest:
    ```bash
    kubectl apply -f ingress-basic.yaml
    ```

2. Get the Ingress address:
    ```bash
    kubectl get ingress
    ```
