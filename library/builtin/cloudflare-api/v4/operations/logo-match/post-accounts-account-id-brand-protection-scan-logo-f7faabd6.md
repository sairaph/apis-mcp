---
title: Create new logo queries from image files
page_id: operation-post-accounts-account-id-brand-protection-scan-logo-2af26ec6
path: operations/logo-match
description: Return new logo queries created from image files
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/scan-logo
operation_ids:
    - postAccountsAccountIdBrandProtectionScanLogo
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new logo queries from image files

`POST /accounts/{account_id}/brand-protection/scan-logo`

Operation ID: `postAccountsAccountIdBrandProtectionScanLogo`

Return new logo queries created from image files

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionScanLogo", "summary": "Create new logo queries from image files", "description": "Return new logo queries created from image files", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "scan-logo", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
