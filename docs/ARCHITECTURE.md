# PulseRoute Architecture & Gateway Specs

PulseRoute sits at the edge of the microservice mesh, managing authentication, sliding rate limits, and reverse proxying.

> [!NOTE]
> High-throughput endpoints must maintain p99 latencies below 1.2ms under peak load.

> [!WARNING]
> The `TokenBucket` rate limiter must be synchronized across concurrent callers to prevent race conditions during bursts.

## Topology Flow

| Tier | Component | Responsibilities | SLA Target |
| :--- | :--- | :--- | :--- |
| **L1** | Cloudflare Edge | DDoS & SSL termination | < 15ms |
| **L2** | PulseRoute Proxy | Token auth & Rate limiting | < 1.5ms |
| **L3** | Service Pods | Business logic execution | < 80ms |

### Routing Diagram
![Architecture Diagram](topology.png)

## Key Modules

- [`cmd/gateway/main.go`](file://cmd/gateway/main.go): HTTP handler setup and routing.
- [`internal/limiter/limiter.go`](file://internal/limiter/limiter.go): In-memory TokenBucket algorithm.
- [`internal/proxy/reverse.go`](file://internal/proxy/reverse.go): Upstream reverse proxy and header stamping.
