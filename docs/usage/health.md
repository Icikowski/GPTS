---
title: Liveness & readiness probes
---

# Accessing **GPTS** liveness & readiness endpoints

Liveness and readiness probe are the way that **GPTS** uses to inform other components about its health. Those are primarily used by Kubernetes, but can be also checked manually.

!!! info "No content returned"
    Both endpoints are optimized to return no content and use only status code to indicate whether the service is running (`200 OK`) or not (`503 Service Unavailable`).

## Liveness endpoint

Liveness endpoint indicates whether application itself (not the exposed service) is up and running.

```bash
# Change ${GPTS} to proper GPTS healthchecks URL (usually exposed on port 8081)
curl ${GPTS}/live
```

???- summary "Example command execution & output"
    ```bash
    curl http://localhost:8081/live -w "${http_code}"
    ```
    ```
    {}
    200
    ```

## Readiness endpoint

Readiness endpoint indicates whether both application and the exposed service are up and running.


```bash
# Change ${GPTS} to proper GPTS healthchecks URL (usually exposed on port 8081)
curl ${GPTS}/ready
```

???- summary "Example command execution & output"
    ```bash
    curl http://localhost:8081/ready -w "${http_code}"
    ```
    ```
    {}
    200
    ```

## Possible statuses

| Application status | Service status | Liveness probe status | Readiness probe status |
|-|-|-|-|
| not ready | not ready | `503 Service Unavailable` | `503 Service Unavailable` |
| ready | not ready | `200 OK` | `503 Service Unavailable` |
| ready | ready | `200 OK` | `200 OK` |

!!! warning "Prohibited combination"
    It is impossible for application to be _not ready_ and service to be _ready_, so such configuration was omitted in the table above.
