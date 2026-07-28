---
title: Create allowlist prefix.
page_id: operation-post-accounts-account-id-magic-advanced-tcp-protection-configs-allowlist-135e9265
path: operations/dos-flowtrackd-api-other
description: Create an allowlist prefix for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist
operation_ids:
    - createAllowlistedPrefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create allowlist prefix.

`POST /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist`

Operation ID: `createAllowlistedPrefix`

Create an allowlist prefix for an account.

## Definition

```yaml
{"operationId": "createAllowlistedPrefix", "summary": "Create allowlist prefix.", "description": "Create an allowlist prefix for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The new allowlist prefix to create.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_NewInfraPrefix"}}}}, "responses": {"200": {"description": "Create allowlist prefix response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_infra-prefix-response"}}}}, "4XX": {"description": "Create allowlist prefix failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
