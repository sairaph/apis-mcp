---
title: Read all brands
page_id: operation-get-accounts-account-id-brand-protection-brands-1e8603f8
path: operations/brand-protection
description: Return all brands
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/brands
operation_ids:
    - getAccountsAccountIdBrandProtectionBrands
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read all brands

`GET /accounts/{account_id}/brand-protection/brands`

Operation ID: `getAccountsAccountIdBrandProtectionBrands`

Return all brands

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionBrands", "summary": "Read all brands", "description": "Return all brands", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "brands", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
