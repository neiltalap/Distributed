# Storage

Managing storage in Kubernetes involves multiple resources:

## Manifests

| File | Description |
|------|-------------|
| `persistent-volume.yaml` | **PersistentVolume (PV)**: A piece of storage in the cluster provisioned by an administrator. |
| `persistent-volume-claim.yaml` | **PersistentVolumeClaim (PVC)**: A request for storage by a user. |
| `storage-class.yaml` | **StorageClass**: Describes the "classes" of storage offered (e.g., SSD, HDD). |
| `pod-with-pvc.yaml` | A Pod that mounts a PVC. |
| `pvc-dynamic.yaml` | A PVC that requests dynamic provisioning via a StorageClass. |
| `pod-emptydir.yaml` | **emptyDir**: Ephemeral storage that lasts as long as the Pod. |
| `pod-hostpath.yaml` | **hostPath**: Mounts a file or directory from the host node's filesystem. |
| `pod-gitrepo.yaml` | (Deprecated) `gitRepo` volume type. |
| `pod-git-sync.yaml` | Sidecar pattern to sync a git repo to a volume (modern alternative to `gitRepo`). |

## Usage

1. **Static Provisioning**:
    ```bash
    kubectl apply -f persistent-volume.yaml
    kubectl apply -f persistent-volume-claim.yaml
    kubectl apply -f pod-with-pvc.yaml
    ```

2. **Dynamic Provisioning**:
    ```bash
    kubectl apply -f storage-class.yaml
    kubectl apply -f pvc-dynamic.yaml
    ```
