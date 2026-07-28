---
title: Run liveness checks
page_id: operation-get-live-b7269231
path: operations/brand-protection
description: Return a success message after running liveness checks
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /live
operation_ids:
    - getLive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run liveness checks

`GET /live`

Operation ID: `getLive`

Return a success message after running liveness checks

## Definition

```yaml
{"operationId": "getLive", "summary": "Run liveness checks", "description": "Return a success message after running liveness checks", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "live", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
