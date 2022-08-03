---
title: Changelog
---

# Project changelog

## GPTS 0.7.7

- Updated Go version to `1.19`
- Upgraded project's dependencies

## GPTS 0.7.6

- Applied code improvements
- Added changelog for Helm chart
- Added chart signing for Helm chart

## GPTS 0.7.5

- Improved logging code in handlers
- Upgraded project's dependencies
    - `CVE-2022-28948` vulnerability in `gopkg.in/yaml.v3` package was fixed

## GPTS 0.7.4

- Updated Go version to `1.18`
- Upgraded project's dependencies
- Major refactoring of code
    - Renamed package `logging` to `logs`
    - Changed log level resolver
    - Changed log writer selection logic
    - Replaced general log instance fetcher with component-specific one
    - Changed log preparation logic across all packages
    - Changed environment variables getters to use generics
    - Changed `interface{}` notation wit `any`
    - Changed pointer-related tools to use generics
    - Changed `CurrentConfiguration` object creation logic
    - Removed unused constants
    - Rephrased some log messages
    - Added missing EOL at the end of the `Dockerfile`
- Updated `Makefile` targets

## GPTS 0.7.3

- Changed liveness and readiness endpoints provider to `github.com/Icikowski/kubeprobes`
- Upgraded project's dependencies
- Updated documentation

## GPTS 0.7.2

- Updated chart definition
    - Added startup probe
    - Changed default values of `failureThreshold` and `periodSeconds` of all probes
    - Added `overrides` section to override name, fullname, image repository, image name & image tag
    - Added reference for GPTS' documentation for charts' values
    - Added metadata to `Chart.yaml`
- Added support for configurable configuration endpoint's address
- Upgraded project's dependencies
- Changed default handler's response structure
    - Headers' values are now returned as single comma-separated string instead of list
    - Added query values logging
- Updated OpenAPI specification
- Updated documentation

## GPTS 0.7.1

- Upgraded project's dependencies

## GPTS 0.7.0

- Changed base image to [Google's "distroless"](https://github.com/GoogleContainerTools/distroless)

## GPTS 0.6.6

- Fixed fetching configuration from environment variables

## GPTS 0.6.5

- Added support for command line flags (configuration may be passed either by CLI flags or environment variables, where flags are prioritized)

## GPTS 0.6.4

- Redesigned logging architecture
- Minor code improvements

## GPTS 0.6.3

- Added static binary compilation for standalone binaries and Docker image
- Enabled build information logging
- Added build information to binaries

## GPTS 0.6.2

- Renamed environment variables in Helm chart
- Renamed environment variables in Code
- Added log level selection to Helm chart
- Changed log levels of particular messages
- Added possibility to control desired log level
- Changed name of service readiness probe

## GPTS 0.6.1

- Fixed `YAML` vs `JSON` field names inconsistency

## GPTS 0.6.0

- Changed default ports in Helm chart
- Disabled loading default config on startup in Helm chart
- Changed Dockerfile to build rootless image

## GPTS 0.5.1

- Fixed AllowSubpaths option support
- Switched to `httptest.Server` from `http.Server` in tests

## GPTS 0.5.0

- Redesigned configuration structure
- Updated OpenAPI specification
- Updated Go version and project dependencies
- Extended tests coverage

## Early development versions

- Prepared Helm chart
- Prepared Dockerfile
- Added default configuration entry (`/hello` route)
- Added subpaths handling feature
- Introduced declarative configuration model
- Added default response for unconfigured routes
- Initial codebase of General Purpose Test Service
