---
title: Update protection status.
page_id: operation-patch-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-prot-a9f2f018
path: operations/dos-flowtrackd-api-other
description: Update the protection status of the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_protection_status
operation_ids:
    - updateProtectionStatus
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update protection status.

`PATCH /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_protection_status`

Operation ID: `updateProtectionStatus`

Update the protection status of the account.

## Definition

```yaml
{"operationId": "updateProtectionStatus", "summary": "Update protection status.", "description": "Update the protection status of the account.", "parameters": [{"name": "account_id", "in": "path", "description": "The account ID.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The update to apply to the protection status.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_UpdateProtectionStatus"}}}}, "responses": {"200": {"description": "Update protection status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_protection-status-response"}}}}, "4XX": {"description": "Update protection status failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
