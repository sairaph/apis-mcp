---
title: Run readiness checks
page_id: operation-get-ready-2147b936
path: operations/brand-protection
description: Return a success message after running readiness checks
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /ready
operation_ids:
    - getReady
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run readiness checks

`GET /ready`

Operation ID: `getReady`

Return a success message after running readiness checks

## Definition

```yaml
{"operationId": "getReady", "summary": "Run readiness checks", "description": "Return a success message after running readiness checks", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ready", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
