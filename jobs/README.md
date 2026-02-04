# Job

A **Job** creates one or more Pods and ensures that a specified number of them successfully terminate. As pods successfully complete, the Job tracks the successful completions.

## Usage

1. Apply the Job manifest:
    ```bash
    kubectl apply -f job.yaml
    ```
    *(Or `multi-completion-job.yaml` / `parallel-job.yaml` for advanced patterns)*

2. Verify the job completion:
    ```bash
    kubectl get jobs
    ```
