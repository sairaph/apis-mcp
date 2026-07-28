---
title: Update verification statuses of submitted URLs to awaiting by ID
page_id: operation-patch-accounts-account-id-brand-protection-clear-1272bcb4
path: operations/brand-protection
description: Return a success message after updating verification statuses of submitted URLs to awaiting by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/brand-protection/clear
operation_ids:
    - patchAccountsAccountIdBrandProtectionClear
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update verification statuses of submitted URLs to awaiting by ID

`PATCH /accounts/{account_id}/brand-protection/clear`

Operation ID: `patchAccountsAccountIdBrandProtectionClear`

Return a success message after updating verification statuses of submitted URLs to awaiting by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "patchAccountsAccountIdBrandProtectionClear", "summary": "Update verification statuses of submitted URLs to awaiting by ID", "description": "Return a success message after updating verification statuses of submitted URLs to awaiting by ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "clear", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
