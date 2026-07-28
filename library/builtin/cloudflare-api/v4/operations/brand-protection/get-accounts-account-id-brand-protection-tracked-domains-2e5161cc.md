---
title: Read submitted domains by pattern
page_id: operation-get-accounts-account-id-brand-protection-tracked-domains-ceaf8605
path: operations/brand-protection
description: Return submitted domains based on pattern
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/tracked-domains
operation_ids:
    - getAccountsAccountIdBrandProtectionTrackedDomains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read submitted domains by pattern

`GET /accounts/{account_id}/brand-protection/tracked-domains`

Operation ID: `getAccountsAccountIdBrandProtectionTrackedDomains`

Return submitted domains based on pattern

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionTrackedDomains", "summary": "Read submitted domains by pattern", "description": "Return submitted domains based on pattern", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "tracked-domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
