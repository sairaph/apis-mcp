---
title: Create new URL submissions
page_id: operation-post-accounts-account-id-brand-protection-submit-f2cf7d76
path: operations/brand-protection
description: Return new URL submissions
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/submit
operation_ids:
    - postAccountsAccountIdBrandProtectionSubmit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new URL submissions

`POST /accounts/{account_id}/brand-protection/submit`

Operation ID: `postAccountsAccountIdBrandProtectionSubmit`

Return new URL submissions

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionSubmit", "summary": "Create new URL submissions", "description": "Return new URL submissions", "responses": {"201": {"description": "Created", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_URLSubmit"}}}}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "submit", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
