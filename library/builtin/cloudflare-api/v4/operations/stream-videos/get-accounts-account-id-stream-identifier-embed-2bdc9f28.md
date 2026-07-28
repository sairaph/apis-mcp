---
title: Retrieve embed Code HTML
page_id: operation-get-accounts-account-id-stream-identifier-embed-e21d3f87
path: operations/stream-videos
description: Fetches an HTML code snippet to embed a video in a web page delivered through Cloudflare. On success, returns an HTML fragment for use on web pages to display a video. On failure, returns a JSON response body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/embed
operation_ids:
    - stream-videos-retreieve-embed-code-html
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve embed Code HTML

`GET /accounts/{account_id}/stream/{identifier}/embed`

Operation ID: `stream-videos-retreieve-embed-code-html`

Fetches an HTML code snippet to embed a video in a web page delivered through Cloudflare. On success, returns an HTML fragment for use on web pages to display a video. On failure, returns a JSON response body.

## Definition

```yaml
{"operationId": "stream-videos-retreieve-embed-code-html", "summary": "Retrieve embed Code HTML", "description": "Fetches an HTML code snippet to embed a video in a web page delivered through Cloudflare. On success, returns an HTML fragment for use on web pages to display a video. On failure, returns a JSON response body.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Retreieve embed Code HTML response.", "content": {"text/html": {"schema": {"type": "string", "example": "<stream id=\"ea95132c15732412d22c1476fa83f27a\"></stream><script data-cfasync=\"false\" defer type=\"text/javascript\" src=\"https://embed.cloudflarestream.com/embed/we4g.fla9.latest.js\"></script>"}}}}, "4XX": {"description": "Retreieve embed Code HTML response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.embed", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
