---
title: Read string queries by ID
page_id: operation-get-accounts-account-id-brand-protection-queries-65ea81a4
path: operations/domain-search
description: Return string queries based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/queries
operation_ids:
    - getAccountsAccountIdBrandProtectionQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read string queries by ID

`GET /accounts/{account_id}/brand-protection/queries`

Operation ID: `getAccountsAccountIdBrandProtectionQueries`

Return string queries based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionQueries", "summary": "Read string queries by ID", "description": "Return string queries based on ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queries", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
