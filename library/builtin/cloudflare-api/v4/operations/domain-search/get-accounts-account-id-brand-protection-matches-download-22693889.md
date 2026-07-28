---
title: Download matches for string queries by ID
page_id: operation-get-accounts-account-id-brand-protection-matches-download-b1d7c7aa
path: operations/domain-search
description: Return matches as CSV for string queries based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/matches/download
operation_ids:
    - getAccountsAccountIdBrandProtectionMatchesDownload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download matches for string queries by ID

`GET /accounts/{account_id}/brand-protection/matches/download`

Operation ID: `getAccountsAccountIdBrandProtectionMatchesDownload`

Return matches as CSV for string queries based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionMatchesDownload", "summary": "Download matches for string queries by ID", "description": "Return matches as CSV for string queries based on ID", "parameters": [{"name": "id", "in": "query", "schema": {"type": "string"}}, {"name": "offset", "in": "query", "schema": {"type": "integer"}}, {"name": "limit", "in": "query", "schema": {"type": "integer"}}, {"name": "include_domain_id", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_QueryMatch"}}}}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "matches.download", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
