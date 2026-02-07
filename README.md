# Distributed Systems Playground

A structured collection of Kubernetes manifests and Go application examples for learning distributed systems.

## 🚀 Quick Start

### Prerequisites
- **Kubernetes Cluster**: Minikube, Kind, K3s, or a cloud provider.
- **kubectl**: Kubernetes command-line tool.
- **Go** (Optional): For building the Go application locally.
- **Docker** (Optional): For building container images.

---

## 📂 Repository Structure

### Core Concepts
- **[Pods](./pods)**: Basic pod definitions (labeled, GPU, etc.).
- **[Cluster Setup](./cluster-setup)**: Namespace and initial cluster configuration.
- **[Concepts](./concepts)**: Theoretical notes (e.g., Linux PSI, [Standard Labels](./concepts/standard-labels.md)).

### Workloads
- **[ReplicaSets](./replicasets)**: Maintains a stable set of replica Pods.
- **[DaemonSets](./daemonsets)**: Runs a copy of a Pod on all (or selected) Nodes.
- **[Jobs](./jobs)**: Run-to-completion tasks.
- **[CronJobs](./cronjobs)**: Scheduled time-based jobs.
- **[Replication Controllers](./replication-controllers)**: (Legacy) Predecessor to ReplicaSets.

### Networking
- **[Services](./services)**: Exposing applications (ClusterIP, NodePort, LoadBalancer).
- **[Ingresses](./ingresses)**: HTTP/HTTPS routing.
- **[Gateway API](./gateway-api)**: Modern, flexible routing (Gateway, HTTPRoute).

### Storage & Config
- **[Storage](./storage)**: Persistent Volumes (PV), Claims (PVC), and StorageClasses.
- **[Configs](./configs)**: Configuration management using ConfigMaps and Secrets.

### Applications
- **[Go App Demo](./go-app-demo)**: A production-ready Go application demonstrating Kubernetes best practices (Probes, Graceful Shutdown).

---

## 🛠 Usage

1. **Deploy the Demo App**:
    ```bash
    make deploy-app
    ```

2. **Explore a Concept**:
    Navigate to a directory and apply the manifests:
    ```bash
    cd jobs
    kubectl apply -f job.yaml
    ```
