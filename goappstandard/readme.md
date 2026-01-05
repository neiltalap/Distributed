# Kubernetes-Native Go Application Standard

This project demonstrates the **three golden rules** of production-ready Go applications in Kubernetes:

## 1. The Three Probes

* **Startup Probe**: Prevents the container from being killed while it performs heavy tasks (like DB migrations) on boot.
* **Liveness Probe**: Restarts the pod only if the process is deadlocked or crashed.
* **Readiness Probe**: Removes the pod from the Service LoadBalancer if the DB is down, but keeps the pod alive so it can recover.

## 2. Signal Handling (SIGTERM)

When K8s stops a pod, it sends a `SIGTERM`. Our code uses `signal.NotifyContext` to catch this. If we didn't, the app would be killed immediately, potentially corrupting the SQLite database or dropping active user requests.

## 3. Graceful Sequencing

Shutdown happens in a specific order:

1. **Stop Traffic**: `app.ShutdownWithContext` stops the HTTP server.
2. **Finish Requests**: Fiber finishes processing any requests already in progress.
3. **Cleanup**: Only after the server is off do we close the **Database** and background routines.

## Development

To run locally:

```bash
DATABASE_URL=./local.db go run .
