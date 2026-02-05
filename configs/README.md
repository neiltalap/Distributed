# Configuration

Kubernetes provides mechanisms to decouple configuration artifacts from image content.

## Manifests

| File | Description |
|------|-------------|
| `configmap-and-secret.yaml` | Defines a **ConfigMap** (non-sensitive) and a **Secret** (sensitive data). |
| `pod-env-vars.yaml` | Injecting ConfigMap values as Environment Variables. |
| `pod-config-volume.yaml` | Mounting ConfigMaps and Secrets as volumes (files). |

## Usage

1. Create the config and secret:
    ```bash
    kubectl apply -f configmap-and-secret.yaml
    ```

2. Consume them in a Pod:
    ```bash
    kubectl apply -f pod-env-vars.yaml
    # OR
    kubectl apply -f pod-config-volume.yaml
    ```
