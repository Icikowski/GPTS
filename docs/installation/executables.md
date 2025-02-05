---
title: Executable files
---

# Running **GPTS** directly from executables

## Using prebuilt binaries

**GPTS** binaries can be downloaded directly from [releases page of project's repository](https://git.sr.ht/~icikowski/gpts/releases).

## Manually building binaries

!!! important "Prerequisites"
    In order to proceed, you must prepare following prerequisites:

    - **Go** language, version **1.19** or later.

### Cloning the repository

First thing you have to do in order to compile application on your own is cloning the [project's repository](https://git.sr.ht/~icikowski/gpts):

```bash
git clone https://git.sr.ht/~icikowski/gpts
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
    go: downloading github.com/rs/zerolog v1.25.0
    go: downloading github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
    go: downloading github.com/gorilla/mux v1.8.0
    go: downloading gopkg.in/yaml.v2 v2.4.0
    go: downloading github.com/prometheus/client_golang v1.11.0
    go: downloading golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0
    go: downloading github.com/prometheus/client_model v0.2.0
    go: downloading github.com/prometheus/common v0.30.0
    go: downloading github.com/golang/protobuf v1.5.2
    go: downloading github.com/prometheus/procfs v0.7.3
    go: downloading github.com/cespare/xxhash/v2 v2.1.2
    go: downloading github.com/beorn7/perks v1.0.1
    go: downloading google.golang.org/protobuf v1.27.1
    go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.1
    ```

### Building binary

The last step is to build binary file. It is done by executing `go build` command:

```bash
go build .
```

Alternatively, you can use `Makefile` target:

```bash
make build
```

By default, compiled binary will be named `gpts` (or `gpts.exe` on Windows), but you can change it, as well as target OS and architecture, accordingly to docs available after executing command:

```bash
go help build
```
