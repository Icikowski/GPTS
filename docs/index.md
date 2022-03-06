---
title: About project
---

# **GPTS** / General Purpose Test Service

![Go version](https://img.shields.io/github/go-mod/go-version/Icikowski/GPTS?filename=application%2Fgo.mod&style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/Icikowski/GPTS?style=for-the-badge)](https://goreportcard.com/report/github.com/Icikowski/GPTS)
![Codecov](https://img.shields.io/codecov/c/gh/icikowski/GPTS?style=for-the-badge&token=FRS94GYIE7)
[![Helm Chart](https://img.shields.io/badge/dynamic/yaml?color=0f1689&label=Helm%20Chart&query=%24.entries.gpts[0].version&url=https%3A%2F%2Fcharts.icikowski.pl%2Findex.yaml&style=for-the-badge)](https://charts.icikowski.pl)
![License](https://img.shields.io/github/license/Icikowski/GPTS?style=for-the-badge)
[![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FIcikowski%2FGPTS%2Fmaster%2Fopenapi.yml&style=for-the-badge)](https://editor.swagger.io/?url=https%3A%2F%2Fraw.githubusercontent.com%2FIcikowski%2FGPTS%2Fmaster%2Fopenapi.yml)

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
