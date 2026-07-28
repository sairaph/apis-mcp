---
title: Delete saved string queries by ID
page_id: operation-delete-accounts-account-id-brand-protection-queries-ab24b2e2
path: operations/domain-search
description: Return a success message after deleting saved string queries by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/brand-protection/queries
operation_ids:
    - deleteAccountsAccountIdBrandProtectionQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete saved string queries by ID

`DELETE /accounts/{account_id}/brand-protection/queries`

Operation ID: `deleteAccountsAccountIdBrandProtectionQueries`

Return a success message after deleting saved string queries by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "deleteAccountsAccountIdBrandProtectionQueries", "summary": "Delete saved string queries by ID", "description": "Return a success message after deleting saved string queries by ID", "parameters": [{"name": "id", "in": "query", "schema": {"type": "string"}}, {"name": "tag", "in": "query", "schema": {"type": "string"}}, {"name": "scan", "in": "query", "schema": {"type": "boolean"}}], "responses": {"204": {"description": "No Content"}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queries", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
