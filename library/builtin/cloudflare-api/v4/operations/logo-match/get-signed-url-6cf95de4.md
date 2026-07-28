---
title: Internal route for testing signed URLs
page_id: operation-get-signed-url-cca9a607
path: operations/logo-match
description: Internal route for testing signed URLs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /signed-url
operation_ids:
    - getSignedUrl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Internal route for testing signed URLs

`GET /signed-url`

Operation ID: `getSignedUrl`

Internal route for testing signed URLs.

## Definition

```yaml
{"operationId": "getSignedUrl", "summary": "Internal route for testing signed URLs", "description": "Internal route for testing signed URLs.", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "signed-url", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
