---
title: Create new saved string queries
page_id: operation-post-accounts-account-id-brand-protection-queries-00d61a93
path: operations/domain-search
description: Return a success message after creating new saved string queries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/queries
operation_ids:
    - postAccountsAccountIdBrandProtectionQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new saved string queries

`POST /accounts/{account_id}/brand-protection/queries`

Operation ID: `postAccountsAccountIdBrandProtectionQueries`

Return a success message after creating new saved string queries

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionQueries", "summary": "Create new saved string queries", "description": "Return a success message after creating new saved string queries", "parameters": [{"name": "id", "in": "query", "schema": {"type": "string"}}, {"name": "tag", "in": "query", "schema": {"type": "string"}}, {"name": "scan", "in": "query", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_Query"}}}}, "responses": {"204": {"description": "No Content"}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queries", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
