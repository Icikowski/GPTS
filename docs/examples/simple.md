# Simple **GPTS** configuration examples

## Single JSON response on every URI

=== "`application/json`"
    ```json
    {
        "/": {
            "allow_subpaths": true,
            "content_type": "application/json",
            "content": "{\"message\":\"Hello World!\"}"
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /:
      allow_subpaths: true
      content_type: application/json
      content: '{"message":"Hello World!"}'
    ```

## Simple HTML page

=== "`application/json`"
    ```json
    {
        "/": {
            "allow_subpaths": true,
            "content_type": "text/html",
            "content": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><title>Homepage</title></head><body><h1>Welcome on my page</h1><ul><li><a href=\"/about\">About me</a></li><li><a href=\"/contact\">Contact</a></li></ul></body></html>"
        },
        "/about": {
            "allow_subpaths": false,
            "content_type": "text/html",
            "content": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><title>About me</title></head><body><h1>About me</h1><p>Hello, my name is <strong>John Doe</strong>!</p><p>Some text, blah, blah, blah...</p></body></html>"
        },
        "/contact": {
            "allow_subpaths": false,
            "content_type": "text/html",
            "content": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><title>Contact</title></head><body><h1>Contact</h1><p><strong>John Doe</strong><ul><li>Phone: 555-0123-456</li><li>E-mail: <a href=\"mailto:test@example.com\">test@example.com</li></ul></p></body></html>"
        }
    }
    ```
=== "`text/yaml`"
    ```yaml
    /:
      allow_subpaths: true
      content_type: text/html
      content: |
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="utf-8" />
            <title>Homepage</title>
          </head>
          <body>
            <h1>Welcome on my page</h1>
            <ul>
              <li><a href="/about">About me</a></li>
              <li><a href="/contact">Contact</a></li>
            </ul>
          </body>
        </html>
    /about:
      allow_subpaths: false
      content_type: text/html
      content: |
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="utf-8" />
            <title>About me</title>
          </head>
          <body>
            <h1>About me</h1>
            <p>Hello, my name is <strong>John Doe</strong>!</p>
            <p>Some text, blah, blah, blah...</p>
          </body>
        </html>
    /contact:
      allow_subpaths: false
      content_type: text/html
      content: |
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="utf-8" />
            <title>Contact</title>
          </head>
          <body>
            <h1>Contact</h1>
            <p>
              <strong>John Doe</strong>
              <ul>
                <li>Phone: 555-0123-456</li>
                <li>E-mail: <a href="mailto:test@example.com">test@example.com</li>
              </ul>
            </p>
          </body>
        </html>
    ```