---
title: Download matches for logo queries by ID
page_id: operation-get-accounts-account-id-brand-protection-logo-matches-download-d13b6e30
path: operations/logo-match
description: Return matches as CSV for logo queries based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/logo-matches/download
operation_ids:
    - getAccountsAccountIdBrandProtectionLogoMatchesDownload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download matches for logo queries by ID

`GET /accounts/{account_id}/brand-protection/logo-matches/download`

Operation ID: `getAccountsAccountIdBrandProtectionLogoMatchesDownload`

Return matches as CSV for logo queries based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionLogoMatchesDownload", "summary": "Download matches for logo queries by ID", "description": "Return matches as CSV for logo queries based on ID", "parameters": [{"name": "logo_id", "in": "query", "schema": {"type": "array", "items": {"type": "string"}}, "explode": true, "style": "form"}, {"name": "offset", "in": "query", "schema": {"type": "string"}}, {"name": "limit", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_LogoMatch"}}}}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logo-matches.download", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
