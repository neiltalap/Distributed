# Gateway API

The **Gateway API** splits networking into two distinct layers:
1. **Infrastructure (Gateway)**: "I am open on port 80."
2. **Routing (Routes)**: "Send traffic from port 80 to Service X."

**You need BOTH.** A Gateway without a Route is like a phone with no contact list—it has a dial tone (open port) but doesn't know who to call.

## Components

- **Gateway**: Defines the functionality of the load balancer (Ports, TLS).
- **HTTPRoute / TCPRoute**: Defines the traffic rules (Path matching, Backend Services).

## Manifests

| File | Role | Description |
|------|------|-------------|
| `gateway-class.yaml` | **Infrastructure** | Defines the controller (e.g., Nginx, Istio) that manages Gateways. |
| `gateway-basic.yaml` | **Infrastructure** | Opens ports (e.g., 80, 5432). **Does not send traffic anywhere by itself.** |
| `route-http.yaml` | **Routing** | Connects the Gateway to a Service. *Also* demonstrates advanced features like header injection (optional). |
| `route-tcp.yaml` | **Routing** | Connects the TCP Gateway to a DB Service. |

## Usage

1. **Apply the Gateway** (Opens the door):
    ```bash
    kubectl apply -f gateway-basic.yaml
    ```

2. **Apply the Route** (Tells traffic where to go):
    ```bash
    kubectl apply -f route-http.yaml
    ```
