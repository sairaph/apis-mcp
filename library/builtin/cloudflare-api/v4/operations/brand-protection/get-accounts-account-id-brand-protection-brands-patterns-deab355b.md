---
title: Read patterns for brands by ID
page_id: operation-get-accounts-account-id-brand-protection-brands-patterns-505c811a
path: operations/brand-protection
description: Return patterns for brands based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/brands/patterns
operation_ids:
    - getAccountsAccountIdBrandProtectionBrandsPatterns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read patterns for brands by ID

`GET /accounts/{account_id}/brand-protection/brands/patterns`

Operation ID: `getAccountsAccountIdBrandProtectionBrandsPatterns`

Return patterns for brands based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionBrandsPatterns", "summary": "Read patterns for brands by ID", "description": "Return patterns for brands based on ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "brands.patterns", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
