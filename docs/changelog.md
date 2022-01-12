# Project changelog

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
