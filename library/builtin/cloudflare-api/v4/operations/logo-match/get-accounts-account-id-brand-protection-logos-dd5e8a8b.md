---
title: Read all saved logo queries
page_id: operation-get-accounts-account-id-brand-protection-logos-c8a30b16
path: operations/logo-match
description: Return all saved logo queries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/logos
operation_ids:
    - getAccountsAccountIdBrandProtectionLogos
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read all saved logo queries

`GET /accounts/{account_id}/brand-protection/logos`

Operation ID: `getAccountsAccountIdBrandProtectionLogos`

Return all saved logo queries

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionLogos", "summary": "Read all saved logo queries", "description": "Return all saved logo queries", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logos", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
