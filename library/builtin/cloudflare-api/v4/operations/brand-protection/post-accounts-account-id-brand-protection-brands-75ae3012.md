---
title: Create new brands
page_id: operation-post-accounts-account-id-brand-protection-brands-28289eb5
path: operations/brand-protection
description: Return new brands
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/brands
operation_ids:
    - postAccountsAccountIdBrandProtectionBrands
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new brands

`POST /accounts/{account_id}/brand-protection/brands`

Operation ID: `postAccountsAccountIdBrandProtectionBrands`

Return new brands

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionBrands", "summary": "Create new brands", "description": "Return new brands", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "brands", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
