# Running **GPTS** directly from executables

## Using prebuilt binaries

**GPTS** binaries can be downloaded directly from [releases page of project's repository](https://github.com/Icikowski/GPTS/releases).

## Manually building binaries

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Go** language, version **1.16** or later

### Cloning the repository

First thing you have to do in order to compile application on your own is cloning rhe project's repository:

=== "Clone via HTTPS"
    ```bash
    git clone https://github.com/Icikowski/GPTS.git
    ```
=== "Clone via SSH"
    ```bash
    git clone git@github.com:Icikowski/GPTS.git
    ```

### Downloading dependencies

After cloning the source files, you should download all of the project's dependencies. You can do it by executing following command inside `application` subdirectory of **GPTS** repository:

```bash
go get ./...
```

???- summary "Example command execution & output"
    ```bash
    go get ./...
    ```
    ```
    go: downloading github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
    go: downloading github.com/gorilla/mux v1.8.0
    go: downloading gopkg.in/yaml.v2 v2.4.0
    go: downloading github.com/prometheus/client_golang v1.10.0
    go: downloading github.com/golang/protobuf v1.5.2
    go: downloading github.com/prometheus/client_model v0.2.0
    go: downloading github.com/prometheus/common v0.21.0
    go: downloading github.com/cespare/xxhash/v2 v2.1.1
    go: downloading golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7
    go: downloading github.com/beorn7/perks v1.0.1
    go: downloading github.com/prometheus/procfs v0.6.0
    go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.1
    go: downloading google.golang.org/protobuf v1.26.0
    ```

### Building binary

The last step is to build binary file. It is done by executing `go build` command:

```bash
go build .
```

By default, compiled binary will be named `gpts`, but you can change it, as well as target OS and architecture, accordingly to docs available after executing command:

```bash
go help build
```
