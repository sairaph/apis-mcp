---
title: Read the total number of saved string queries
page_id: operation-get-accounts-account-id-brand-protection-total-queries-881e95c2
path: operations/domain-search
description: Return the total number of saved string queries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/total-queries
operation_ids:
    - getAccountsAccountIdBrandProtectionTotalQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read the total number of saved string queries

`GET /accounts/{account_id}/brand-protection/total-queries`

Operation ID: `getAccountsAccountIdBrandProtectionTotalQueries`

Return the total number of saved string queries

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionTotalQueries", "summary": "Read the total number of saved string queries", "description": "Return the total number of saved string queries", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "total-queries", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
