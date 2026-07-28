---
title: Update verification statuses of submitted URLs to confirmed by ID
page_id: operation-patch-accounts-account-id-brand-protection-verify-7e8a7007
path: operations/brand-protection
description: Return a success message after updating verification statuses of submitted URLs to confirmed by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/brand-protection/verify
operation_ids:
    - patchAccountsAccountIdBrandProtectionVerify
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update verification statuses of submitted URLs to confirmed by ID

`PATCH /accounts/{account_id}/brand-protection/verify`

Operation ID: `patchAccountsAccountIdBrandProtectionVerify`

Return a success message after updating verification statuses of submitted URLs to confirmed by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "patchAccountsAccountIdBrandProtectionVerify", "summary": "Update verification statuses of submitted URLs to confirmed by ID", "description": "Return a success message after updating verification statuses of submitted URLs to confirmed by ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "verify", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
