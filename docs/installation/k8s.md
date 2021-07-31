# Deploying **GPTS** on Kubernetes cluster via Helm

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Kubernetes cluster** (tested with Kubernetes bundled with [_Docker Desktop_](https://www.docker.com/products/docker-desktop));
    - **Helm 3** (although there's a slight possibility that Helm 2 will also work);
    - **Ingress controller** (_optional_; tested with [_NGINX Ingress Controller_](https://kubernetes.github.io/ingress-nginx/)).

## Configuring repository

In order to obtain Helm chart of **GPTS**, you need to configure proper repository. You can do it by executing following commands:

```bash
helm repo add icikowski https://charts.icikowski.pl
helm repo update
```

## Fetching chart from repository

After repository is successfully added, you can check for available versions of `gpts` chart:

```bash
helm search repo gpts
```

???- summary "Example command output"
    ```
    NAME            CHART VERSION   APP VERSION     DESCRIPTION
    icikowski/gpts  0.1.0           0.1.0           GPTS - General Purpose Test Service
    ```

In order to fetch chart, execute one of following commands:

=== "Fetch latest as .tgz"
    ```bash
    helm fetch icikowski/gpts
    ```
=== "Fetch latest and unpack it"
    ```bash
    helm fetch icikowski/gpts --untar
    ```
=== "Fetch particular version as .tgz"
    ```bash
    # For example: chart version 0.1.0
    helm fetch icikowski/gpts:0.1.0
    ```
=== "Fetch particular version and unpack it"
    ```bash
    # For example: chart version 0.1.0
    helm fetch icikowski/gpts:0.1.0 --untar
    ```

!!! info "Downloading chart directly"
    It is also possible to download charts manually [from chart repository](https://charts.icikowski.pl), eg. for offline use.

## Changing values in chart

!!! warning "Development phase"
    Project is in early development phase. Section will be expanded after release.

    Three environment variables are used to determine GPTS behavior:
    
    - **GPTS_SERVER_PORT** to specify service port (default: "80")
    - **GPTS_HEALTHCHECK_PORT** to specify healthcheck endpoints port (default: "8000")
    - **GPTS_DEFAULT_CONFIG_ON_STARTUP** to specify whether default config (/hello endpoint) should be loaded when application starts (default: "false")

    Those values can be configured with following variables in `values.yaml`:

    ```yaml
    gpts:
      servicePort: 80
      healthcheckPort: 8000
      defaultConfigOnStartup: true
    ```

In order to configure chart before deployment (eg. enable ingress, change service type), you need to change values in `values.yaml` file inside chart's directory.

???- example "Example: enabling ingress for NGINX ingress class"
    === "Default values"
        ```yaml linenums="30" hl_lines="2 3 4 7"
        ingress:
          enabled: false
          annotations: {}
            # kubernetes.io/ingress.class: nginx
            # kubernetes.io/tls-acme: "true"
          hosts:
          - host: chart-example.local
            paths:
            - path: /
          tls: []
          #  - secretName: chart-example-tls
          #    hosts:
          #    - chart-example.local
        ```
    === "Modified values"
        ```yaml linenums="30" hl_lines="2 3 4 7"
        ingress:
          enabled: true
          annotations:
            kubernetes.io/ingress.class: nginx
            # kubernetes.io/tls-acme: "true"
          hosts:
          - host: test0.host.net
            paths:
            - path: /
          tls: []
          #  - secretName: chart-example-tls
          #    hosts:
          #    - chart-example.local
        ```

## Deploying chart on cluster

As the chart is prepared, you can install it on your cluster. First of all, let's prepare a namespace for deployment:

```bash
kubectl create ns NAMESPACE
```

???- summary "Example command execution & output"
    ```bash
    kubectl create ns test-service
    ```
    ```
    namespace/test-service created
    ```

The last step is the actual deployment:

```bash
helm install -n NAMESPACE DEPLOYMENT_NAME CHART_DIRECTORY
```

???- summary "Example command execution & output"
    ```bash
    helm install -n test-service my-service ./gpts
    ```
    ```
    NAME: my-service
    LAST DEPLOYED: Mon Apr 26 21:56:34 2021
    NAMESPACE: test-service
    STATUS: deployed
    REVISION: 1
    NOTES:
    1. Get the application URL by running these commands:
       http://test0.host.net/
    ```

Application is up and running now! You can check it by cURLing the ingress address (if you enabled it in `values.yaml`).

???- summary "Example command execution & output"
    ```bash
    curl -s http://test0.host.net | jq
    ```
    ```json
    {
        "host": "test0.host.net",
        "path": "/",
        "method": "GET",
        "headers": {
            "Accept": [
                "*/*"
            ],
            "User-Agent": [
                "curl/7.68.0"
            ],
            "X-Forwarded-For": [
                "192.168.65.3"
            ],
            "X-Forwarded-Host": [
                "test0.host.net"
            ],
            "X-Forwarded-Port": [
                "80"
            ],
            "X-Forwarded-Proto": [
                "http"
            ],
            "X-Real-Ip": [
                "192.168.65.3"
            ],
            "X-Request-Id": [
                "b73fb52f9ee980ea9e5218993434b535"
            ],
            "X-Scheme": [
                "http"
            ]
        }
    }
    ```

!!! note "Next steps: configuring endpoints"
    In order to configure endpoints, please check out the [_Configuring endpoints_ section of _User guide_](../usage/endpoints.md).
