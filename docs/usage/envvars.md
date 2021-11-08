# Controlling **GPTS** setting with environment variables

There are four configurable **GPTS** settings which can be controlled using listed environment variables:

- `GPTS_SERVICE_PORT`,
- `GPTS_HEALTHCHECKS_PORT`,
- `GPTS_DEFAULT_CONFIG_ON_STARTUP`,
- `GPTS_LOG_LEVEL`,
- `GPTS_PRETTY_LOG`.

## `GPTS_SERVICE_PORT`

**Description**: Defines the port on which service will be exposed; defaults to `80` in most cases, but is set to `8080` in K8s deployments due to rootless mode[^1] limitations.

[^1]: **Rootless mode** is related to privileges level. **GPTS** images are built in the way that ensures usage of lowest possible privileges (user `app`, UID `1000`) in order to maintain proper security level. Due to lack of `root` privileges, the default service port (`80`) cannot be bind properly, therefore it was changed to `8080`.

**Allowed values**: valid port number

**Default value**: `80` (or `8080` in Helm charts)

## `GPTS_HEALTHCHECKS_PORT`

**Description**: Defines the port on which health endpoints will be exposed.

**Allowed values**: valid port number

**Default value**: `8081`

## `GPTS_DEFAULT_CONFIG_ON_STARTUP`

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

## `GPTS_LOG_LEVEL`

**Description**: Defines the desired log level in accordance to [`zerolog`'s log levels](https://github.com/rs/zerolog#leveled-logging).

**Allowed values**: `debug`, `info`, `warn`, `error`, `fatal`, `panic`, `trace`

**Default value**: `info`

## `GPTS_PRETTY_LOG`

**Description**: Enables/Disables pretty log format (may impact performance)

**Allowed values**: `false` or `true`

**Default value**: `false`
