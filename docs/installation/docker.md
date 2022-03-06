---
title: Docker
---

# Deploying **GPTS** in Docker

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Docker** engine available;
    - Internet connection available.

!!! info "Determining the latest version tag"
    Since the images are not tagged with `latest` tag, you can manually check the latest version by visiting the address [https://ghcr.io/icikowski/gpts](https://ghcr.io/icikowski/gpts).

!!! info "Changing default settings"
    **GPTS** settings are configured with environment variables. [Read more](../usage/flags.md) about those variables and change configuration according to your needs by using `-e`/`--env` flag in `docker` command or using `environment` section of `docker-compose.yml` file.

## Using **docker run** command

You can use one of following commands in order to get the application up and running:

=== "Expose only test service"
    ```bash
    # Replace ${SERVICE_PORT} with port you want the service to be exposed on 
    # Replace ${TAG} with desired image tag
    docker run --rm -it -p ${SERVICE_PORT}:80 ghcr.io/icikowski/gpts:${TAG}
    ```

    ???- summary "Example command execution & output"
        ```bash
        docker run --rm -it -p 80:80 ghcr.io/icikowski/gpts:0.6.2
        ```
        ```
        0.6.2: Pulling from icikowski/gpts
        a0d0a0d46f8b: Already exists
        de2ec1b71fb1: Pull complete
        d5a5180019b8: Pull complete
        7f926084d672: Pull complete
        Digest: sha256:446ede875a67ab19196d5dd30905171dc87fd333392319c5073e15555d5a6974
        Status: Downloaded newer image for ghcr.io/icikowski/gpts:0.6.2
        {"level":"info","component":"cli","servicePort":"80","healthchecksPort":"8081","time":"2021-10-01T15:33:26Z","message":"starting application"}
        {"level":"info","component":"service","time":"2021-10-01T15:33:26Z","message":"preparing test service's router & server"}
        {"level":"info","component":"service","time":"2021-10-01T15:33:26Z","message":"preparing configuration handler"}
        {"level":"info","component":"service","time":"2021-10-01T15:33:26Z","message":"preparing default handler"}
        {"level":"info","component":"service","time":"2021-10-01T15:33:26Z","message":"server prepared"}
        ```
=== "Expose both test service and liveness & readiness probes"
    ```bash
    # Replace ${SERVICE_PORT} with port you want the service to be exposed on 
    # Replace ${HEALTHCHECKS_PORT} with port you want the application's health endpoints to be exposed on
    # Replace ${TAG} with desired image tag
    docker run --rm -it -p ${SERVICE_PORT}:80 -p ${HEALTHCHECKS_PORT}:8081 ghcr.io/icikowski/gpts:${TAG}
    ```
    
    ???- summary "Example command execution & output"
        ```bash
        docker run --rm -it -p 80:80 -p 8000:8000 ghcr.io/icikowski/gpts:0.6.2
        ```
        ```
        Unable to find image 'ghcr.io/icikowski/gpts:0.6.2' locally
        0.6.2: Pulling from icikowski/gpts
        a0d0a0d46f8b: Already exists
        de2ec1b71fb1: Pull complete
        d5a5180019b8: Pull complete
        7f926084d672: Pull complete
        Digest: sha256:446ede875a67ab19196d5dd30905171dc87fd333392319c5073e15555d5a6974
        Status: Downloaded newer image for ghcr.io/icikowski/gpts:0.6.2
        {"level":"info","component":"cli","servicePort":"80","healthchecksPort":"8081","time":"2021-10-01T15:35:57Z","message":"starting application"}
        {"level":"info","component":"service","time":"2021-10-01T15:35:57Z","message":"preparing test service's router & server"}
        {"level":"info","component":"service","time":"2021-10-01T15:35:57Z","message":"preparing configuration handler"}
        {"level":"info","component":"service","time":"2021-10-01T15:35:57Z","message":"preparing default handler"}
        {"level":"info","component":"service","time":"2021-10-01T15:35:57Z","message":"server prepared"}
        ```

## Using **docker compose** command

In order to use `docker compose` or `docker-compose` deployment method, you'll have to create `docker-compose.yml` file as follows:

```yaml
version: "3.3"

# Replace ${TAG} with desired image tag
# Replace ${SERVICE_PORT} with port you want the service to be exposed on 
# Replace ${HEALTHCHECKS_PORT} with port you want the application's health endpoints to be exposed on

services:
  gpts:
    image: ghcr.io/icikowski/gpts:${TAG}
    ports:
      - "${SERVICE_PORT}:80"
      - "${HEALTHCHECKS_PORT}:8081" # This one is optional, you can remove this line
    # environment:
    #  - GPTS_SERVICE_PORT=80
    #  - GPTS_HEALTHCHECKS_PORT=8081
    #  - GPTS_CONFIG_ENDPOINT=/config
    #  - GPTS_DEFAULT_CONFIG_ON_STARTUP=false
    #  - GPTS_LOG_LEVEL=info
    #  - GPTS_PRETTY_LOG=false

```

???- example "Example contents of docker-compose.yml"
    ```yaml
    version: "3.3"

    services:
      gpts:
        image: ghcr.io/icikowski/gpts:0.6.4
        ports:
          - "80:80"
          - "8081:8081" # This one is optional, you can remove this line
        environment:
          - GPTS_DEFAULT_CONFIG_ON_STARTUP=true
          - GPTS_LOG_LEVEL=debug
          - GPTS_PRETTY_LOG=true
    ```

After the file is saved, you can deploy the application by executing `docker compose up -d` or `docker-compose up -d` command in directory which contains the file.

???- summary "Example command execution & output"
    ```bash
    docker compose up -d
    ```
    ```
    [+] Running 5/5
    - gpts Pulled                                                     3.4s
    - a0d0a0d46f8b Already exists                                   0.0s
    - de2ec1b71fb1 Pull complete                                    0.9s
    - d5a5180019b8 Pull complete                                    0.9s
    - 7f926084d672 Pull complete                                    1.7s
    [+] Running 2/2
    - Network gpts_default   Created                                  0.0s
    - Container gpts_gpts_1  Started                                  0.9s
    ```
