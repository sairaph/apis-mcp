---
title: Create new saved string queries in bulk
page_id: operation-post-accounts-account-id-brand-protection-queries-bulk-a3172945
path: operations/domain-search
description: Return a success message after creating new saved string queries in bulk
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/queries/bulk
operation_ids:
    - postAccountsAccountIdBrandProtectionQueriesBulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new saved string queries in bulk

`POST /accounts/{account_id}/brand-protection/queries/bulk`

Operation ID: `postAccountsAccountIdBrandProtectionQueriesBulk`

Return a success message after creating new saved string queries in bulk

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionQueriesBulk", "summary": "Create new saved string queries in bulk", "description": "Return a success message after creating new saved string queries in bulk", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_QueryBulk"}}}}, "responses": {"204": {"description": "No Content"}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queries.bulk", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
