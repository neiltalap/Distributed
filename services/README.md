# Services

A **Service** is an abstract way to expose an application running on a set of Pods as a network service.
![NodePort](nodeport.png)

## Manifests

| File | Description |
|------|-------------|
| `service-clusterip.yaml` | **ClusterIP** (Default): Exposes the Service on a cluster-internal IP. Reachable only within the cluster. |
| `service-nodeport.yaml` | **NodePort**: Exposes the Service on each Node's IP at a static port (30000-32767). |
| `service-loadbalancer.yaml` | **LoadBalancer**: Exposes the Service externally using a cloud provider's load balancer. |
| `service-named-ports.yaml` | Demonstrates mapping ports by name rather than number. |
| `service-multi-port.yaml` | Exposes multiple ports (e.g., HTTP and HTTPS) on a single Service. |

## Usage

### ClusterIP
```bash
kubectl apply -f service-clusterip.yaml
kubectl get svc my-service
```

### NodePort
```bash
kubectl apply -f service-nodeport.yaml
# Access via <NodeIP>:30080
```

### LoadBalancer
```bash
kubectl apply -f service-loadbalancer.yaml
# External IP will be provisioned by cloud provider (or Minikube tunnel)
```
