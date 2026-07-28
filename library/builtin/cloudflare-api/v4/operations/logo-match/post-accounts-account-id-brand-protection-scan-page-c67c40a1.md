---
title: Create new logo queries from URLs
page_id: operation-post-accounts-account-id-brand-protection-scan-page-3903bf98
path: operations/logo-match
description: Return new logo queries created from URLs
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/scan-page
operation_ids:
    - postAccountsAccountIdBrandProtectionScanPage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new logo queries from URLs

`POST /accounts/{account_id}/brand-protection/scan-page`

Operation ID: `postAccountsAccountIdBrandProtectionScanPage`

Return new logo queries created from URLs

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionScanPage", "summary": "Create new logo queries from URLs", "description": "Return new logo queries created from URLs", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "scan-page", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
