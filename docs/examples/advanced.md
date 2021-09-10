# Advanced **GPTS** configuration examples

## HTML page with images

=== "`application/json`"
    ```json
    {
        "/": {
            "allow_subpaths": true,
            "content_type": "text/html",
            "content": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><link rel=\"icon\" href=\"/favicon.png\" /><title>Hello, World</title></head><body><h1>Hello, world!</h1><img src=\"/smile.jpg\" /></body></html>"
        },
        "/favicon.png": {
            "allow_subpaths": false,
            "content_type": "image/png",
            "content": "base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAACJSURBVDhP7ZPbDcAgCACdi4GYx2lYhl202mCrrc9omn54P4oxFxBUdhFb1OZbkTFGdmWKIiatiVmiNmWRBiUAYFtaEyG5xWeGEKT+KE9LFEO4RQl/ErFvt2/4lMglIiMzKXKXA6DDLBOOD+SV0f1CkXtAdMLxv6oVG5EVPZDKXl1M6BF1sUhk7QEtmishFSnrxwAAAABJRU5ErkJggg=="
        },
        "/smile.jpg": {
            "allow_subpaths": false,
            "content_type": "image/jpeg",
            "content": "base64,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"
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /:
      allow_subpaths: true
      content_type: text/html
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
      allow_subpaths: false
      content_type: image/png
      content: |-
        base64,
        iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAAXNSR0IArs4c6QAAAARnQU1BAACx
        jwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAACJSURBVDhP7ZPbDcAgCACdi4GYx2lYhl202mCr
        rc9omn54P4oxFxBUdhFb1OZbkTFGdmWKIiatiVmiNmWRBiUAYFtaEyG5xWeGEKT+KE9LFEO4RQl/
        ErFvt2/4lMglIiMzKXKXA6DDLBOOD+SV0f1CkXtAdMLxv6oVG5EVPZDKXl1M6BF1sUhk7QEtmish
        FSnrxwAAAABJRU5ErkJggg==
    /smile.jpg:
      allow_subpaths: false
      content_type: image/jpeg
      content: |-
        base64,
        /9j/4AAQSkZJRgABAQEAeAB4AAD/2wBDAAIBAQIBAQICAgICAgICAwUDAwMDAwYEBAMFBwYHBwcG
        BwcICQsJCAgKCAcHCg0KCgsMDAwMBwkODw0MDgsMDAz/2wBDAQICAgMDAwYDAwYMCAcIDAwMDAwM
        DAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAz/wAARCACAAIADASIA
        AhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQA
        AAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3
        ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWm
        p6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEA
        AwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSEx
        BhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElK
        U1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3
        uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD9AKKK
        K/yPPnwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA
        KKKKACiiigAooooAKKKKACuH/aI/aU8Cfsm/Dd/FvxG8T6d4T0LMkdvPd7y9/MiBzb28aK0k0pBX
        5UU4DBm2rlhV/as/af8AC37GP7PviH4l+M/tR0Lw/GFW3tsCbU7yQMLazjYghXmddu7DbFEkhUrG
        1eC/8E+f+Cauv/taeK7D9pj9rS1/4SDxzqyrceDfAd2jDRvAenbi8Km1csDM2Q4jfOwsZJfMuXYx
        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
    ```

## Custom statuses and headers

=== "`application/json`"
    ```json
    {
        "/error": {
            "allow_subpaths": true,
            "status": 400,
            "content_type": "application/json",
            "content": "{\"message\": \"Bad Request!\"}"
        },
        "/error/403": {
            "allow_subpaths": true,
            "status": 403,
            "content_type": "application/json",
            "content": "{\"message\": \"Unauthorized!\"}"
        },
        "/error/404": {
            "allow_subpaths": true,
            "status": 404,
            "content_type": "application/json",
            "content": "{\"message\": \"Not found!\"}"
        },
        "/error/500": {
            "allow_subpaths": true,
            "status": 500,
            "content_type": "application/json",
            "content": "{\"message\": \"Internal Server Error!\"}"
        },
        "/redirect-to-google": {
            "allow_subpaths": true,
            "status": 307,
            "headers": {
            "Location": "https://google.com"
            }
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /error:
      allow_subpaths: true
      status: 400
      content_type: application/json
      content: |-
        {
          "message": "Bad Request!"
        }
    /error/403:
      allow_subpaths: true
      status: 403
      content_type: application/json
      content: |-
        {
          "message": "Unauthorized!"
        }
    /error/404:
      allow_subpaths: true
      status: 404
      content_type: application/json
      content: |-
        {
          "message": "Not found!"
        }
    /error/500:
      allow_subpaths: true
      status: 500
      content_type: application/json
      content: |-
        {
          "message": "Internal Server Error!"
        }
    /redirect-to-google:
      allow_subpaths: true
      status: 307
      headers:
        Location: https://google.com
    ```
