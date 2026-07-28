---
title: Create new string queries
page_id: operation-post-accounts-account-id-brand-protection-search-e7bc31b0
path: operations/domain-search
description: Return new string queries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/search
operation_ids:
    - postAccountsAccountIdBrandProtectionSearch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new string queries

`POST /accounts/{account_id}/brand-protection/search`

Operation ID: `postAccountsAccountIdBrandProtectionSearch`

Return new string queries

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionSearch", "summary": "Create new string queries", "description": "Return new string queries", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "search", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
