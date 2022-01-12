# **GPTS**' default response for unconfigured routes

If user is attempting to send a request for unconfigured route, the application will respond with a predefined response, which consists of following fields:

| Field name | Value type | Description |
|-|-|-|
| `host` | `string` | Hostname for which the request was sent |
| `path` | `string` | Path for which the request was sent |
| `method` | `string` | HTTP method of the sent request |
| `headers` | `string` to `string` map | HTTP headers of the sent request |
| `queries` | `string` to `string` map | Query values of the sent request |

By default the response is returned in JSON format, but it is possible to retrieve it in YAML format by using `Accept: text/yaml` header.

=== "Example of JSON response"
    ```json
    {
      "host": "localhost:8080",
      "path": "/some-path",
      "method": "GET",
      "headers": {
        "Accept": "*/*",
        "User-Agent": "curl/7.81.0"
      },
      "queries": {
        "param1": "someValue",
        "param2": "someOtherValue"
      }
    }
    ```
=== "Example of YAML response"
    ```yaml
    host: localhost:8080
    path: /
    method: GET
    headers:
      Accept: text/yaml
      User-Agent: curl/7.81.0
    queries:
      param1: someValue
      param2: someOtherValue
    ```

!!! important "Changing the default response"
    Please note that default response is hardcoded in the application, so it cannot be changed directly. In order to override it, the route for path `/` with `allowSubpaths` option enabled must be manually created and added to the configuration.
