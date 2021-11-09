# Deploying **GPTS** on Kubernetes cluster via Helm

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Kubernetes cluster** (tested with Kubernetes bundled with [_Docker Desktop_](https://www.docker.com/products/docker-desktop));
    - **Helm 3** (although there's a slight possibility that Helm 2 will also work);
    - **Ingress controller** (_optional_; tested with [_NGINX Ingress Controller_](https://kubernetes.github.io/ingress-nginx/));
    - Internet connection available.

## Configuring repository

In order to obtain Helm chart of **GPTS**, you need to configure proper repository. You can do it by executing following commands:

```bash
helm repo add icikowski https://charts.icikowski.pl
helm repo update
```

## Fetching chart from repository

After repository is successfully added, you can check for available versions of `gpts` chart:

=== "Latest version"
    ```bash
    helm search repo gpts
    ```
    ???- summary "Example command output"
        ```
        NAME            CHART VERSION   APP VERSION     DESCRIPTION
        icikowski/gpts  0.6.2           0.6.2           GPTS - General Purpose Test Service
        ```
=== "All versions"
    ```bash
    helm search repo gpts -l
    ```
    ???- summary "Example command output"
        ```
        NAME            CHART VERSION   APP VERSION     DESCRIPTION
        icikowski/gpts  0.6.2           0.6.2           GPTS - General Purpose Test Service
        icikowski/gpts  0.6.1           0.6.1           GPTS - General Purpose Test Service
        icikowski/gpts  0.6.0           0.6.0           GPTS - General Purpose Test Service
        icikowski/gpts  0.5.1           0.5.1           GPTS - General Purpose Test Service
        icikowski/gpts  0.5.0           0.5.0           GPTS - General Purpose Test Service
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
    # For example: chart version 0.6.0
    helm fetch icikowski/gpts --version 0.6.0
    ```
=== "Fetch particular version and unpack it"
    ```bash
    # For example: chart version 0.6.0
    helm fetch icikowski/gpts --version 0.6.0 --untar
    ```

!!! info "Downloading chart directly"
    It is also possible to download charts manually [from chart repository](https://charts.icikowski.pl), eg. for offline use.

## Changing configuration values in chart

**GPTS** settings are configured with environment variables [described here](../usage/flags.md) and can be set using following values:

| Chart value | Environment variable | Default value |
|-|-|-|
| `gpts.servicePort` | `GPTS_SERVICE_PORT` | `8080` |
| `gpts.healthchecksPort` | `GPTS_HEALTHCHECKS_PORT` | `8081` |
| `gpts.defaultConfigOnStartup` | `GPTS_DEFAULT_CONFIG_ON_STARTUP` | `false` |
| `gpts.logLevel` | `GPTS_LOG_LEVEL` | `info` |
| `gpts.prettyLog` | `GPTS_PRETTY_LOG` | `false` |

???- example "Example contents of _gpts_ section in values.yaml"
    ```yaml linenums="11"
    gpts:
      servicePort: 8080
      healthchecksPort: 8081
      defaultConfigOnStartup: false
      logLevel: info
      # Available log levels: 
      # debug, info, warn, error, fatal, panic, trace

      # Enabling pretty log can make the logs more user-friendly
      # but is NOT RECOMMENDED as it impacts the performance a lot
      prettyLog: false
    ```

In order to configure chart before deployment (eg. enable ingress, change service type), you need to change values in `values.yaml` file inside chart's directory.

???- example "Example: enabling ingress for NGINX ingress class"
    === "Default values"
        ```yaml linenums="23" hl_lines="2 3 4 7"
        ingress:
          enabled: false
          annotations: {}
            # kubernetes.io/ingress.class: nginx
            # kubernetes.io/tls-acme: "true"
          hosts:
          - host: example.com
            paths:
            - path: /
          tls: []
          #  - secretName: chart-example-tls
          #    hosts:
          #    - example.com
        ```
    === "Modified values"
        ```yaml linenums="23" hl_lines="2 3 4 7"
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
    LAST DEPLOYED: Fri Oct  1 18:13:34 2021
    NAMESPACE: test-service
    STATUS: deployed
    REVISION: 1
    NOTES:
    GPTS - General Purpose Test Service

    Service installed successfully! Check out the documentation (https://icikowski.github.io/GPTS) and start using the app.

    Get the application URL by running these commands:
      http://test0.host.net/
    ```

Application is up and running now! You can check it by cURLing the ingress address (if you enabled it in `values.yaml`).

???- summary "Example command execution & output (JSON format; default)"
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
                "curl/7.79.1"
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
            "X-Forwarded-Scheme": [
                "http"
            ],
            "X-Real-Ip": [
                "192.168.65.3"
            ],
            "X-Request-Id": [
                "b08e79506d17797a0f890fea0579978e"
            ],
            "X-Scheme": [
                "http"
            ]
        }
    }
    ```

???- summary "Example command execution & output (YAML format)"
    ```bash
    curl -s http://test0.host.net -H "Accept: text/yaml" | yq eval -
    ```
    ```yaml
    host: test0.host.net
    path: /
    method: GET
    headers:
      Accept:
        - text/yaml
      User-Agent:
        - curl/7.79.1
      X-Forwarded-For:
        - 192.168.65.3
      X-Forwarded-Host:
        - test0.host.net
      X-Forwarded-Port:
        - "80"
      X-Forwarded-Proto:
        - http
      X-Forwarded-Scheme:
        - http
      X-Real-Ip:
        - 192.168.65.3
      X-Request-Id:
        - be588ef84375e7207cdf8c1c9acfe731
      X-Scheme:
        - http
    ```
