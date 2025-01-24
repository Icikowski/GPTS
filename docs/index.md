---
title: About project
---

# **GPTS** / General Purpose Test Service

Simple & easy to configure test service for serving content through HTTP.

## Features

- [X] Default echo response fon unconfigured routes
- [X] Declarative configuration
    - [X] Support `YAML` and `JSON` configuration
    - [X] Support per-method response definitions (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`)
    - [X] Support default response definition (for unconfigured methods)
    - [X] Support subpaths handling
- [X] Docker support
    - [X] Based on latest [Google's "distroless"](https://github.com/GoogleContainerTools/distroless) image
    - [X] Small size (< 20MB)
    - [X] Running in rootless mode
- [X] Kubernetes support
    - [X] Helm chart available
    - [X] Image running as non-root
    - [X] All settings can be configured via chart values
    - [X] Support Ingress controllers
- [ ] [More coming soon...](devplans.md)

## Development model

**GPTS** is being developed as a hobby project. Due to large workload caused by ongoing studies and professional work, I might not be able to spend a lot of time on further improvements and new features. I'll still try to fix any bugs and vulnerabilities as soon as possible, but general development will be conducted during holidays and/or after finishing current term.
