---
title: Read submitted domains by ID
page_id: operation-get-accounts-account-id-brand-protection-domain-info-1e911bf3
path: operations/brand-protection
description: Return submitted domains based on ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/brand-protection/domain-info
operation_ids:
    - getAccountsAccountIdBrandProtectionDomainInfo
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read submitted domains by ID

`GET /accounts/{account_id}/brand-protection/domain-info`

Operation ID: `getAccountsAccountIdBrandProtectionDomainInfo`

Return submitted domains based on ID

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "getAccountsAccountIdBrandProtectionDomainInfo", "summary": "Read submitted domains by ID", "description": "Return submitted domains based on ID", "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["brand_protection"], "x-api-token-group": ["Intel Write", "Intel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "domain-info", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
