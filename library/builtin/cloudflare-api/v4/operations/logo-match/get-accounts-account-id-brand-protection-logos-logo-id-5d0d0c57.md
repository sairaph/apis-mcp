---
title: Read saved logo queries by ID
page_id: operation-get-accounts-account-id-brand-protection-logos-logo-id-8b1fd393
path: operations/logo-match
description: Return saved logo queries based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/logos/{logo_id}
operation_ids:
    - getAccountsAccountIdBrandProtectionLogosLogoId
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read saved logo queries by ID

`GET /accounts/{account_id}/brand-protection/logos/{logo_id}`

Operation ID: `getAccountsAccountIdBrandProtectionLogosLogoId`

Return saved logo queries based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "logo_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionLogosLogoId", "summary": "Read saved logo queries by ID", "description": "Return saved logo queries based on ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logos", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
