---
title: Read all alerts on submitted domains
page_id: operation-get-accounts-account-id-brand-protection-alerts-95525a63
path: operations/brand-protection
description: Return all alerts on submitted domains
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/alerts
operation_ids:
    - getAccountsAccountIdBrandProtectionAlerts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read all alerts on submitted domains

`GET /accounts/{account_id}/brand-protection/alerts`

Operation ID: `getAccountsAccountIdBrandProtectionAlerts`

Return all alerts on submitted domains

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionAlerts", "summary": "Read all alerts on submitted domains", "description": "Return all alerts on submitted domains", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "alerts", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
