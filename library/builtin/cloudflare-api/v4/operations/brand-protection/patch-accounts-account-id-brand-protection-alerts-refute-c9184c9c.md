---
title: Update verification statuses of tracked URLs to disproven by ID
page_id: operation-patch-accounts-account-id-brand-protection-alerts-refute-159861b9
path: operations/brand-protection
description: Return a success message after updating verification statuses of tracked URLs to disproven by ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/brand-protection/alerts/refute
operation_ids:
    - patchAccountsAccountIdBrandProtectionAlertsRefute
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update verification statuses of tracked URLs to disproven by ID

`PATCH /accounts/{account_id}/brand-protection/alerts/refute`

Operation ID: `patchAccountsAccountIdBrandProtectionAlertsRefute`

Return a success message after updating verification statuses of tracked URLs to disproven by ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "patchAccountsAccountIdBrandProtectionAlertsRefute", "summary": "Update verification statuses of tracked URLs to disproven by ID", "description": "Return a success message after updating verification statuses of tracked URLs to disproven by ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "alerts.refute", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
