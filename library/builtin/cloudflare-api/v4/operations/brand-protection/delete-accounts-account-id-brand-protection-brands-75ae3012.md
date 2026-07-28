---
title: Delete brands by ID
page_id: operation-delete-accounts-account-id-brand-protection-brands-3f1e2cec
path: operations/brand-protection
description: Return a success message after deleting brands by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/brand-protection/brands
operation_ids:
    - deleteAccountsAccountIdBrandProtectionBrands
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete brands by ID

`DELETE /accounts/{account_id}/brand-protection/brands`

Operation ID: `deleteAccountsAccountIdBrandProtectionBrands`

Return a success message after deleting brands by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "deleteAccountsAccountIdBrandProtectionBrands", "summary": "Delete brands by ID", "description": "Return a success message after deleting brands by ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "brands", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
