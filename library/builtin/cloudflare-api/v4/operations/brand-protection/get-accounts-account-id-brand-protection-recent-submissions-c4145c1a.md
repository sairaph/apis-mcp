---
title: Read recent URL submissions
page_id: operation-get-accounts-account-id-brand-protection-recent-submissions-99586ad9
path: operations/brand-protection
description: Return recent URL submissions
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/recent-submissions
operation_ids:
    - getAccountsAccountIdBrandProtectionRecentSubmissions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read recent URL submissions

`GET /accounts/{account_id}/brand-protection/recent-submissions`

Operation ID: `getAccountsAccountIdBrandProtectionRecentSubmissions`

Return recent URL submissions

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionRecentSubmissions", "summary": "Read recent URL submissions", "description": "Return recent URL submissions", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "recent-submissions", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
