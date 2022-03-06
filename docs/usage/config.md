---
title: Configuring endpoints
---

# Configuring **GPTS** endpoints

**GPTS** uses declarative configuration which can be fetched or applied using `/config` endpoint.

!!! important "Supported formats"
    **GPTS** supports both JSON and YAML configuration formats, but uses JSON as default. In order to fetch/apply YAML configuration, the proper request header must be set:

    - `Accept: text/yaml` for `GET` requests
    - `Content-Type: text/yaml` for `POST` requests

## Configuration file structure

Declarative configuration used by **GPTS** is a map of URL paths (starting with `/`) and `Route` object, which consist of 7 fields:

| Field name | Description | Allowed values | Default value |
|-|-|-|-|
| `allowSubpaths` | Causes route to be used for all requests made to subpaths of defined URL path (e.g. if base path is `/abc`, then subpaths like `/abc/def`, `/abc/xyx` will be served using this configuration as well) | `true`, `false` | `false` |
| `default` | Default response (used when the specialized response for given method is not available) | `Response` object | `null` |
| `get` | Specialized response for GET request | `Response` object | `null` |
| `post` | Specialized response for POST request | `Response` object | `null` |
| `put` | Specialized response for PUT request | `Response` object | `null` |
| `patch` | Specialized response for PATCH request | `Response` object | `null` |
| `delete` | Specialized response for DELETE request | `Response` object | `null` |

`Response` object defines the details of expected response and consists of 4 fields:

| Field name | Description | Allowed values | Default value |
|-|-|-|-|
| `status` | HTTP status code | [correct HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) as integer | `200` |
| `contentType` | Value of "Content-Type" header; MIME type of returned content | valid MIME type | `text/plain` |
| `content` | Text content or `base64`-encoded binary content (in such case should start with `base64,`) | string | _empty string_ |
| `headers` | Dictionary of additional headers that should be sent in response | string to string map | `null` |

If the configuration is correct, **GPTS** will respond with `202 Accepted` status and no content. In such case, service will be reloaded with new configuration after approximately 2 seconds.

!!! warning "Missing configuration details"
    If neither default nor method-specific response is defined for given route, **GPTS** will respond with `503 Service Unavailable` status and no content.

!!! warning "Incorrect configuration"
    In case of trying to apply malformed configuration, **GPTS**'s `/config` endpoint will respond with following status codes and appropriate message:
    
    - `415 Unsupported Media Type` when sending content different from `application/json` or `text/yaml`
    - `400 Bad Request` if `JSON`/`YAML` content can't be parsed

!!! warning "Other errors"
    If any error would occur as the request is being processed (e.g. when decoding `base64`-encoded content), **GPTS** will respond with `500 Internal Server Error` status and no content.

!!! info "OpenAPI specification"
    For more formalized form of the description above, please look at [OpenAPI specification](https://editor.swagger.io/?url=https%3A%2F%2Fraw.githubusercontent.com%2FIcikowski%2FGPTS%2Fmaster%2Fopenapi.yml) of configuration endpoint.


## Fetching current configuration

In order to fetch current running configuration, `GET` request has to be made for `/config` endpoint.

=== "JSON"
    ```bash
    # Change ${GPTS} to proper GPTS service URL
    curl ${GPTS}/config
    ```

    ???- summary "Example command execution & output"
        ```bash
        curl -s http://localhost/config | jq
        ```
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
    ```bash
    # Change ${GPTS} to proper GPTS service URL
    curl ${GPTS}/config -H "Accept: text/yaml"
    ```
    
    ???- summary "Example command execution & output"
        ```bash
        curl -s http://localhost/config -H "Accept: text/yaml" | yq eval -
        ```
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

## Applying new configuration

In order to apply new configuration, `POST` request has to be made for `/config` endpoint with proper request body.

=== "JSON"
    ```bash
    curl -X POST http://localhost/config -H 'Content-Type: application/json' --data-binary @- <<EOF
    {
        // Content goes here...
    }
    EOF
    ```

    ???- summary "Example command execution & output"
        ```bash
        curl -X POST http://localhost/config -H 'Content-Type: application/json' -w "%{http_code}" --data-binary @- <<EOF
        {
          "/test": {
            "default": {
              "content": "this is some test content"
            }
          }
        }
        EOF
        ```
        ```
        202
        ```

=== "YAML"
    ```bash
    curl -X POST http://localhost/config -H 'Content-Type: text/yaml' --data-binary @- <<EOF
    # Content goes here
    EOF
    ```

    ???- summary "Example command execution & output"
        ```bash
        curl -X POST http://localhost/config -H 'Content-Type: text/yaml' -w "%{http_code}" --data-binary @- <<EOF
        /test:
          default:
            content: this is some test content
        EOF
        ```
        ```
        202
        ```
