---
title: Read URL submissions by ID
page_id: operation-get-accounts-account-id-brand-protection-submission-info-f9b1e767
path: operations/brand-protection
description: Return URL submissions based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/submission-info
operation_ids:
    - getAccountsAccountIdBrandProtectionSubmissionInfo
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read URL submissions by ID

`GET /accounts/{account_id}/brand-protection/submission-info`

Operation ID: `getAccountsAccountIdBrandProtectionSubmissionInfo`

Return URL submissions based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionSubmissionInfo", "summary": "Read URL submissions by ID", "description": "Return URL submissions based on ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "submission-info", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
