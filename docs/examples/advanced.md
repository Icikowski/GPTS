# Advanced **GPTS** configuration examples

## HTML page with images

=== "`application/json`"
    ```json
    {
        "/": {
            "allowSubpaths": true,
            "default": {
              "contentType": "text/html",
              "content": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><link rel=\"icon\" href=\"/favicon.png\" /><title>Hello, World</title></head><body><h1>Hello, world!</h1><img src=\"/smile.jpg\" /></body></html>"
            }
        },
        "/favicon.png": {
            "allowSubpaths": false,
            "default": {
              "contentType": "image/png",
              "content": "base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAACJSURBVDhP7ZPbDcAgCACdi4GYx2lYhl202mCrrc9omn54P4oxFxBUdhFb1OZbkTFGdmWKIiatiVmiNmWRBiUAYFtaEyG5xWeGEKT+KE9LFEO4RQl/ErFvt2/4lMglIiMzKXKXA6DDLBOOD+SV0f1CkXtAdMLxv6oVG5EVPZDKXl1M6BF1sUhk7QEtmishFSnrxwAAAABJRU5ErkJggg=="
            }
        },
        "/smile.jpg": {
            "allowSubpaths": false,
            "default": {
              "contentType": "image/jpeg",
              "content": "base64,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"
            }
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /:
      allowSubpaths: true
      default:
        contentType: text/html
        content: |-
          <!DOCTYPE html>
          <html lang="en">
          <head>
            <meta charset="utf-8" />
            <link rel="icon" href="/favicon.png" />
            <title>Hello, World</title>
          </head>
          <body>
            <h1>Hello, world!</h1>
            <img src="/smile.jpg" />
          </body>
          </html>
    /favicon.png:
      allowSubpaths: false
      default:
        contentType: image/png
        content: |-
          base64,
          iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAAXNSR0IArs4c6QAAAARnQU1BAACx
          jwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAACJSURBVDhP7ZPbDcAgCACdi4GYx2lYhl202mCr
          rc9omn54P4oxFxBUdhFb1OZbkTFGdmWKIiatiVmiNmWRBiUAYFtaEyG5xWeGEKT+KE9LFEO4RQl/
          ErFvt2/4lMglIiMzKXKXA6DDLBOOD+SV0f1CkXtAdMLxv6oVG5EVPZDKXl1M6BF1sUhk7QEtmish
          FSnrxwAAAABJRU5ErkJggg==
    /smile.jpg:
      allowSubpaths: false
      default:
        contentType: image/jpeg
        content: |-
          base64,
          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
          RK5CYII=
    ```

## Custom statuses and headers

=== "`application/json`"
    ```json
    {
        "/error": {
            "allowSubpaths": true,
            "default": {
              "status": 400,
              "contentType": "application/json",
              "content": "{\"message\": \"Bad Request!\"}"
            }
        },
        "/error/403": {
            "allowSubpaths": true,
            "default": {
              "status": 403,
              "contentType": "application/json",
              "content": "{\"message\": \"Unauthorized!\"}"
            }
        },
        "/error/404": {
            "allowSubpaths": true,
            "default": {
              "status": 404,
              "contentType": "application/json",
              "content": "{\"message\": \"Not found!\"}"
            }
        },
        "/error/500": {
            "allowSubpaths": true,
            "default": {
              "status": 500,
              "contentType": "application/json",
              "content": "{\"message\": \"Internal Server Error!\"}"
            }
        },
        "/redirect-to-google": {
            "allowSubpaths": true,
            "default": {
              "status": 307,
              "headers": {
                "Location": "https://google.com"
              }
            }
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /error:
      allowSubpaths: true
      default:
        status: 400
        contentType: application/json
        content: |-
          {
            "message": "Bad Request!"
          }
    /error/403:
      allowSubpaths: true
      default:
        status: 403
        contentType: application/json
        content: |-
          {
            "message": "Unauthorized!"
          }
    /error/404:
      allowSubpaths: true
      default:
        status: 404
        contentType: application/json
        content: |-
          {
            "message": "Not found!"
          }
    /error/500:
      allowSubpaths: true
      default:
        status: 500
        contentType: application/json
        content: |-
          {
            "message": "Internal Server Error!"
          }
    /redirect-to-google:
      allowSubpaths: true
      default:
        status: 307
        headers:
          Location: https://google.com
    ```

## Method-specific responses

=== "`application/json`"
    ```json
    {
      "/endpoint": {
        "allowSubpaths": true,
        "get": {
          "status": 200,
          "contentType": "application/json",
          "content": "{\"message\": \"OK\"}"
        },
        "post": {
          "status": 201,
          "contentType": "application/json",
          "content": "{\"message\": \"Created\"}"
        },
        "put": {
          "status": 409,
          "contentType": "application/json",
          "content": "{\"message\": \"Conflict\"}"
        },
        "patch": {
          "status": 202,
          "contentType": "application/json",
          "content": "{\"message\": \"Accepted\"}"
        },
        "delete": {
          "status": 204,
          "contentType": "application/json",
          "content": "{\"message\": \"No Content\"}"
        }
      }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /endpoint:
      allowSubpaths: true
      get:
        status: 200
        contentType: application/json
        content: |-
          {
            "message": "OK"
          }
      post:
        status: 201
        contentType: application/json
        content: |-
          {
            "message": "Created"
          }
      put:
        status: 409
        contentType: application/json
        content: |-
          {
            "message": "Conflict"
          }
      patch:
        status: 202
        contentType: application/json
        content: |-
          {
            "message": "Accepted"
          }
      delete:
        status: 204
        contentType: application/json
        content: |- 
          {
            "message": "No Content"
          }
    ```
