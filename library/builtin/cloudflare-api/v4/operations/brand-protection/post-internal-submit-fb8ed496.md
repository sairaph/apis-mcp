---
title: Internal route for testing URL submissions
page_id: operation-post-internal-submit-d490ed69
path: operations/brand-protection
description: Internal route for testing URL submissions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /internal/submit
operation_ids:
    - postInternalSubmit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Internal route for testing URL submissions

`POST /internal/submit`

Operation ID: `postInternalSubmit`

Internal route for testing URL submissions.

## Definition

```yaml
{"operationId": "postInternalSubmit", "summary": "Internal route for testing URL submissions", "description": "Internal route for testing URL submissions.", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "internal.submit", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
