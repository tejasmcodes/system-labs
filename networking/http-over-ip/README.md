# HTTP over IP

Experiments to understand what happens when HTTP is built directly
over IP without TCP.

## Experiment 1 — Raw IP Socket

Successfully created a raw IPv4 socket using Linux's raw socket API.

- Address family: IPv4 (`AF_INET`)
- Socket type: `SOCK_RAW`
- Protocol: experimental protocol `253`
- TCP: not used
- UDP: not used
- `CAP_NET_RAW`: required

### Observation

The Linux kernel allows the process to access raw IP networking,
but privileges are required to create the socket.