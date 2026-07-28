---
title: Delete saved logo queries by ID
page_id: operation-delete-accounts-account-id-brand-protection-logos-logo-id-e514dd69
path: operations/logo-match
description: Return a success message after deleting saved logo queries by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/brand-protection/logos/{logo_id}
operation_ids:
    - deleteAccountsAccountIdBrandProtectionLogosLogoId
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete saved logo queries by ID

`DELETE /accounts/{account_id}/brand-protection/logos/{logo_id}`

Operation ID: `deleteAccountsAccountIdBrandProtectionLogosLogoId`

Return a success message after deleting saved logo queries by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "logo_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "deleteAccountsAccountIdBrandProtectionLogosLogoId", "summary": "Delete saved logo queries by ID", "description": "Return a success message after deleting saved logo queries by ID", "responses": {"204": {"description": "No Content"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logos", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
