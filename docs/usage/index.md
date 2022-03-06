---
title: User guide
---

# User guide

**GPTS** has relatively simple configuration, which can be controlled using [flags and environment variables](flags.md).

The process of [fetching and applying new configuration](config.md) consists of sending proper `GET` and `POST` requests to always available `/config` endpoint. In case of sending request for unconfigured route, the [default response](default.md) will be used.

Moreover, it is possible to access [liveness & readiness probes](health.md) in order to check whether the application is working correctly.
