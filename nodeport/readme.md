# Network hops

When an external client connects to a service throught the nodeport.
If it hits the node where the pod is not running an additional hop
is required.

You can prevent this additional hop by configuring the service to redirect external
traffic only to pods running on the node that received the connection. This is done by
setting the externalTrafficPolicy field in the service’s spec section:

```yaml
spec:
 externalTrafficPolicy: Local
 ...
```

![alt text](image.png)

## BEING AWARE OF THE NON-PRESERVATION OF THE CLIENT’S IP

Usually, when clients inside the cluster connect to a service, the pods backing the service can obtain the client’s IP address. But when the connection is received through a node port, the packets’ source IP is changed, because Source Network Address Translation (SNAT) is performed on the packets.

The backing pod can’t see the actual client’s IP, which may be a problem for some
applications that need to know the client’s IP. In the case of a web server, for example, this means the access log won’t show the browser’s IP. The Local external traffic policy described in the previous section affects the preservation of the client’s IP, because there’s no additional hop between the node receiving the connection and the node hosting the target pod (SNAT isn’t performed).
