# DaemonSet

A **DaemonSet** ensures that all (or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected.

## Usage

1. Apply the DaemonSet manifest:
    ```bash
    kubectl apply -f ssd-monitor.yaml
    ```
    *(Note: This example uses a node selector for `disk=ssd`. You may need to label your nodes first using `kubectl label node <node-name> disk=ssd`)*

2. Verify the pods are running on the appropriate nodes:
    ```bash
    kubectl get ds
    kubectl get pods -o wide
    ```
