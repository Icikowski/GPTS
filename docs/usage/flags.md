# Controlling **GPTS** setting with CLI flags/environment variables

There are four configurable **GPTS** settings which can be controlled using listed flags and/or environment variables:

| Option name | Flag | Environment variable |
|-|-|-|
| [Service port](#service-port) | `--service-port` | `GPTS_SERVICE_PORT` |
| [Healthchecks port](#healthchecks-port) | `--healthchecks-port` | `GPTS_HEALTHCHECKS_PORT` |
| [Default configuration on startup](#default-configuration-on-startup)  | `--default-config` | `GPTS_DEFAULT_CONFIG_ON_STARTUP` |
| [Log level](#log-level) | `--log-level` | `GPTS_LOG_LEVEL` |
| [Pretty logging](#pretty-logging) | `--pretty-log` | `GPTS_PRETTY_LOG` |


## Service port

**Description**: Defines the port on which service will be exposed; defaults to `80` in most cases, but is set to `8080` in K8s deployments due to rootless mode[^1] limitations.

[^1]: **Rootless mode** is related to privileges level. **GPTS** images are built in the way that ensures usage of lowest possible privileges (user `app`, UID `1000`) in order to maintain proper security level. Due to lack of `root` privileges, the default service port (`80`) cannot be bind properly, therefore it was changed to `8080`.

**Allowed values**: valid port number

**Default value**: `80` (or `8080` in Helm charts)

## Healthchecks port

**Description**: Defines the port on which health endpoints will be exposed.

**Allowed values**: valid port number

**Default value**: `8081`

## Default configuration on startup

**Description**: Determines whether the default configuration (see below) should be loaded or not on application start.

**Allowed values**: `false` or `true`

**Default value**: `false`

???- example "Default configuration contents"
    === "JSON"
        ```json
        {
          "/hello": {
            "allowSubpaths": true,
            "default": {
              "status": 200,
              "contentType": "application/json",
              "content": "{\"message\":\"Hello World!\"}",
              "headers": {
                "X-SentBy": "GPTS - General Purpose Test Service"
              }
            }
          }
        }    
        ```
    === "YAML"
        ```yaml
        /hello:
          allowSubpaths: true
          default:
            status: 200
            contentType: application/json
            content: '{"message":"Hello World!"}' 
            headers:
              X-SentBy: GPTS - General Purpose Test Service
        ```

## Log level

**Description**: Defines the desired log level in accordance to [`zerolog`'s log levels](https://github.com/rs/zerolog#leveled-logging).

**Allowed values**: `debug`, `info`, `warn`, `error`, `fatal`, `panic`, `trace`

**Default value**: `info`

## Pretty logging

**Description**: Enables/Disables pretty log format (may impact performance)

**Allowed values**: `false` or `true`

**Default value**: `false`
