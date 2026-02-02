# Standard Kubernetes Labels

You might have noticed labels like `app.kubernetes.io/name` or `app.kubernetes.io/part-of` in the manifests. These are not random; they are the **official Kubernetes recommended labels**.

## Why use them?

1.  **Querying**: It allows you to easily filter resources.
    ```bash
    # "Find all resources that are part of the distributed-playground"
    kubectl get all -l app.kubernetes.io/part-of=distributed-playground
    ```
2.  **Tooling**: Many tools (like Helm, Kustomize, ArgoCD, or Dashboards) use these labels to visualize the relationship between objects.
3.  **Consistency**: It avoids the "wild west" of everyone inventing their own labels (`app=backend`, `service=backend`, `tier=backend`...).

## The 3 Key Labels We Use

| Label | Description | Example |
|-------|-------------|---------|
| `app.kubernetes.io/name` | The name of the application. | `go-app-demo`, `kubia` |
| `app.kubernetes.io/component` | The component within the architecture. | `database`, `server`, `worker` |
| `app.kubernetes.io/part-of` | The name of the higher-level application. | `distributed-playground` |

## Example

```yaml
metadata:
  labels:
    app.kubernetes.io/name: mysql
    app.kubernetes.io/component: database
    app.kubernetes.io/part-of: wordpress
```

## Labels vs. `metadata.name`

It's easy to confuse the two, but they have very different purposes:

| Field | Purpose | Analogy | Unique? |
|-------|---------|---------|---------|
| `metadata.name` | **Unique Identifier**. How you talk to *this specific object*. | Social Security Number | **YES** (per namespace) |
| `labels` | **Grouping/Tagging**. How you talk to *groups of objects*. | "Software Engineer" | **NO** (Many pods can have the same label) |

**Example:**
If you have 3 replicas of the `kubia` app:
- They ALL have the label `app.kubernetes.io/name: kubia`.
- But they have DIFFERENT names: `kubia-xc91`, `kubia-aa22`, `kubia-bb33`.

The **Service** uses the *label* to find all 3 of them. It doesn't care about their individual names.
