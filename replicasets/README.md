# ReplicaSet

A **ReplicaSet**'s purpose is to maintain a stable set of replica Pods running at any given time. It is often used to guarantee the availability of a specified number of identical Pods.

## Usage

1. Apply the ReplicaSet manifest:
    ```bash
    kubectl apply -f replicaset.yaml
    ```

2. Verify the replicas are running:
    ```bash
    kubectl get rs
    kubectl get pods
    ```

3. Try deleting a pod and watch it being recreated:
    ```bash
    kubectl delete pod <pod-name>
    ```
