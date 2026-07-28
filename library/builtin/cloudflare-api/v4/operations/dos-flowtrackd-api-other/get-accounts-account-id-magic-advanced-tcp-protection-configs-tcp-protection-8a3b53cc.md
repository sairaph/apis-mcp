---
title: Get protection status.
page_id: operation-get-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-protec-9e5ffbdf
path: operations/dos-flowtrackd-api-other
description: Get the protection status of the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_protection_status
operation_ids:
    - getProtectionStatus
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get protection status.

`GET /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_protection_status`

Operation ID: `getProtectionStatus`

Get the protection status of the account.

## Definition

```yaml
{"operationId": "getProtectionStatus", "summary": "Get protection status.", "description": "Get the protection status of the account.", "parameters": [{"name": "account_id", "in": "path", "description": "The account ID.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Get protection status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_protection-status-response"}}}}, "4XX": {"description": "Get protection status failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
