# Kubernetes Storage Reference

## emptyDir

Type: Ephemeral (Temporary).

Lifespan: Dies with the Pod.

Use Case: Scratch space, cache, sidecar data sharing.

Note: Can use RAM (medium: Memory) for speed.

## gitRepo (Deprecated)

Type: Ephemeral (emptyDir + git clone).

Lifespan: Dies with the Pod.

Use Case: Static websites.

Flaw: Static. No auto-sync. (Use a sidecar container instead).

## hostPath

Type: Persistent (Node-Local).

Lifespan: Lives on the specific Node's disk.

Use Case: System agents needing Node access (e.g., logging).

Warning: Do not use for Apps/DBs. If Pod moves to a new Node, data is lost.

## Direct Cloud Volume (AWS/GCE/NFS)

Type: Persistent (Networked).

Lifespan: Durable. Follows Pod across Nodes.

Flaw: Hardcoded. YAML lists specific Cloud IDs (AWS EBS, etc.), making it non-portable.

## PV & PVC (Manual Provisioning)

Decouples Hardware (Admin) from Request (Developer).

PersistentVolume (PV): The Asset. Admin manually defines the specific storage resource.

PersistentVolumeClaim (PVC): The Ticket. Developer requests size/access.

Binding: Kubernetes locks a matching PV to the PVC.

Critical Settings:

RWO: 1 Node writes. | ROX: Many read. | RWX: Many write.

Reclaim Policies: Retain (Manual cleanup), Recycle (Scrub data), Delete (Destroy volume).

## StorageClass (Dynamic Provisioning)

Automates the creation of PVs. Replaces the "Manual Admin" step.

StorageClass (SC): The Vending Machine. Admin defines how to create storage (e.g., "fast-ssd" using the AWS Provisioner).

Provisioner: The underlying plugin code (CSI Driver) that talks to the cloud API to create the disk.

Workflow:

Developer creates PVC asking for storageClassName: fast-ssd.

Provisioner automatically creates the physical disk.

Provisioner automatically creates the PV object in Kubernetes.

PVC binds to the new PV.

## CSI (Container Storage Interface)

The "Plug and Play" Standard.

Concept: A standardized interface that allows any storage vendor to talk to Kubernetes.

Implementation: You do not write this code.

Vendor (AWS, NetApp, etc.): Writes the CSI Driver (Go code) to translate Kubernetes commands into hardware actions.

It is either the standard implementation k8s has baked into it for AWS and major providers, or just a container that runs in your cluster.

User (You): Simply installs the driver (usually a Pod) provided by the vendor.

Benefit: Decouples Kubernetes releases from storage vendor updates. You can update your storage driver without upgrading the whole cluster.

## Rook (The Operator Pattern)

Software Defined Storage for the cluster.

Concept: Turns local node disks into a unified, smart storage pool (usually Ceph).

Capabilities: Provides Block, File, and Object (S3) storage simultaneously.

Use Case: Running production-grade storage on bare metal (On-Prem) or avoiding cloud vendor lock-in.

Power: It automates complex storage tasks (rebalancing, healing) using the K8s Operator pattern.
