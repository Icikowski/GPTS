# Deploying **GPTS** in Docker

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Docker** engine available

!!! info "Determining the latest version tag"
    Since the images are not tagged with `latest` tag, you can manually check the latest version by visiting the address [https://ghcr.io/icikowski/gpts](https://ghcr.io/icikowski/gpts).

!!! warning "Development phase"
    Project is in early development phase. Section will be expanded after release.
    
    Three environment variables are used to determine GPTS behavior:
    
    - **GPTS_SERVER_PORT** to specify service port (default: "80")
    - **GPTS_HEALTHCHECK_PORT** to specify healthcheck endpoints port (default: "8000")
    - **GPTS_DEFAULT_CONFIG_ON_STARTUP** to specify whether default config (/hello endpoint) should be loaded when application starts (default: "false")


## Using **docker run** command

!!! warning "Development phase"
    Project is in early development phase. Section will be expanded after release.

You can use one of following commands in order to get the application up and running:

=== "Expose only test service"
    ```bash
    # Replace TS_PORT with port you want the test service to be exposed on 
    # Replace TAG with desired image tag
    docker run --rm -it -p TS_PORT:80 ghcr.io/icikowski/gpts:TAG
    ```

    ???- summary "Example command execution & output"
        ```bash
        docker run --rm -it -p 80:80 ghcr.io/icikowski/gpts:0.1.0
        ```
        ```
        Unable to find image 'ghcr.io/icikowski/gpts:0.1.0' locally
        0.1.0: Pulling from icikowski/gpts
        540db60ca938: Already exists
        9323189f6695: Already exists
        4f4fb700ef54: Already exists
        8f7a184cac58: Already exists
        Digest: sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6fe8
        Status: Downloaded newer image for ghcr.io/icikowski/gpts:0.1.0
        2021/04/27 20:16:31 Starting application
        2021/04/27 20:16:31 Preparing readiness & liveness endpoints
        2021/04/27 20:16:31 Loaded default config as current
        2021/04/27 20:16:31 Marking application liveness as UP
        2021/04/27 20:16:31 Launching test service server
        2021/04/27 20:16:31 Preparing test service's router & server
        2021/04/27 20:16:31 Preparing configuration endpoint handler
        2021/04/27 20:16:31 Preparing [/hello] endpoint handler
        2021/04/27 20:16:31 Starting readiness & liveness endpoints at port 8000
        2021/04/27 20:16:31 Preparing default handler
        2021/04/27 20:16:31 Marking test service readiness as UP
        2021/04/27 20:16:31 Starting actual HTTP server
        ```
=== "Expose both test service and liveness & readiness probes"
    ```bash
    # Replace TS_PORT with port you want the test service to be exposed on 
    # Replace LR_PORT with port you want the application's health endpoints to be exposed on
    # Replace TAG with desired image tag
    docker run --rm -it -p TS_PORT:80 -p LR_PORT:8000 ghcr.io/icikowski/gpts:TAG
    ```
    
    ???- summary "Example command execution & output"
        ```bash
        docker run --rm -it -p 80:80 -p 8000:8000 ghcr.io/icikowski/gpts:0.1.0
        ```
        ```
        0.1.0: Pulling from icikowski/gpts
        540db60ca938: Already exists
        9323189f6695: Already exists
        4f4fb700ef54: Already exists
        8f7a184cac58: Already exists
        Digest: sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6fe8
        Status: Downloaded newer image for ghcr.io/icikowski/gpts:0.1.0
        2021/04/27 20:21:51 Starting application
        2021/04/27 20:21:51 Preparing readiness & liveness endpoints
        2021/04/27 20:21:51 Loaded default config as current
        2021/04/27 20:21:51 Marking application liveness as UP
        2021/04/27 20:21:51 Launching test service server
        2021/04/27 20:21:51 Preparing test service's router & server
        2021/04/27 20:21:51 Preparing configuration endpoint handler
        2021/04/27 20:21:51 Starting readiness & liveness endpoints at port 8000
        2021/04/27 20:21:51 Preparing [/hello] endpoint handler
        2021/04/27 20:21:51 Preparing default handler
        2021/04/27 20:21:51 Marking test service readiness as UP
        2021/04/27 20:21:51 Starting actual HTTP server
        ```

## Using **docker-compose**/**docker compose** command

!!! warning "Development phase"
    Project is in early development phase. Section will be expanded after release.

In order to use `docker-compose` or `docker compose` deployment method, you'll have to create `docker-compose.yml` file as follows:

```yaml
version: "3.3"

# Replace TAG with desired image tag
# Replace TS_PORT with port you want the test service to be exposed on 
# Replace LR_PORT with port you want the application's health endpoints to be exposed on

services:
    gpts:
        image: ghcr.io/icikowski/gpts:TAG
        ports:
            - "TS_PORT:80"
            - "LR_PORTS:8000" # This one is optional, you can remove this line

```

???- example "Example contents of docker-compose.yml"
    ```yaml
    version: "3.3"

    services:
        gpts:
            image: ghcr.io/icikowski/gpts:TAG
            ports:
                - "80:80"
                - "8000:8000"
    ```

After the file is saved, you can deploy the application by executing one of following commands in directory which contains the file:

=== "`docker-compose` command"
    ```bash
    docker-compose up -d
    ```

    ???- summary "Example command execution & output"
        ```bash
        docker-compose up -d
        ```
        ```
        Docker Compose is now in the Docker CLI, try `docker compose up`

        Creating network "test_default" with the default driver
        Pulling gpts (ghcr.io/icikowski/gpts:0.1.0)...
        0.1.0: Pulling from icikowski/gpts
        540db60ca938: Already exists
        9323189f6695: Already exists
        4f4fb700ef54: Already exists
        8f7a184cac58: Already exists
        Digest: sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6fe8
        Status: Downloaded newer image for ghcr.io/icikowski/gpts:0.1.0
        Creating test_gpts_1 ... done
        ```

=== "`docker compose` command"
    ```bash
    docker compose up -d
    ```

    ???- summary "Example command execution & output"
        ```bash
        docker compose up -d
        ```
        ```
        [+] Building 2.5s (5/5) FINISHED
         => [internal] load build definition from Dockerfile                                                               0.0s
         => => transferring dockerfile: 70B                                                                                0.0s
         => [internal] load .dockerignore                                                                                  0.0s
         => => transferring context: 2B                                                                                    0.0s
         => [internal] load metadata for ghcr.io/icikowski/gpts:0.1.0                                                      2.1s
         => [1/1] FROM ghcr.io/icikowski/gpts:0.1.0@sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6f  0.2s
         => => resolve ghcr.io/icikowski/gpts:0.1.0@sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6f  0.0s
         => => sha256:fab53b88f0528dbc5ffac55d1b838392ba9577c7756fbe4dcec767d20fbfd4ca 2.08kB / 2.08kB                     0.0s
         => => sha256:2a1b09510588c8e69b73f5196a4fd67bc83ca56aae204a3b1caa254d128a6fe8 1.16kB / 1.16kB                     0.0s
         => exporting to image                                                                                             0.0s
         => => exporting layers                                                                                            0.0s
         => => writing image sha256:fab53b88f0528dbc5ffac55d1b838392ba9577c7756fbe4dcec767d20fbfd4ca                       0.0s
         => => naming to ghcr.io/icikowski/gpts:0.1.0                                                                      0.0s
        [+] Running 2/2
         - Network "test_default"  Created                                                                                 0.1s
         - Container test_gpts_1   Started                                                                                 0.9s
        ```

!!! note "Next steps: configuring endpoints"
    In order to configure endpoints, please check out the [_Configuring endpoints_ section of _User guide_](../usage/endpoints.md).