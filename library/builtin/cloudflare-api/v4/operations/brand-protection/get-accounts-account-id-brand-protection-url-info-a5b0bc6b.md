---
title: Read submitted URLs by ID
page_id: operation-get-accounts-account-id-brand-protection-url-info-7e50a22b
path: operations/brand-protection
description: Return submitted URLs based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/url-info
operation_ids:
    - getAccountsAccountIdBrandProtectionUrlInfo
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read submitted URLs by ID

`GET /accounts/{account_id}/brand-protection/url-info`

Operation ID: `getAccountsAccountIdBrandProtectionUrlInfo`

Return submitted URLs based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionUrlInfo", "summary": "Read submitted URLs by ID", "description": "Return submitted URLs based on ID", "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_URLInfo"}}}}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "url-info", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
