# CronJob

A **CronJob** creates Jobs on a repeating schedule.

## Usage

1. Apply the CronJob manifest:
    ```bash
    kubectl apply -f cronjob.yaml
    ```

2. Watch the jobs being created:
    ```bash
    kubectl get jobs --watch
    ```

3. Check the logs of a completed job pod to see the output.
