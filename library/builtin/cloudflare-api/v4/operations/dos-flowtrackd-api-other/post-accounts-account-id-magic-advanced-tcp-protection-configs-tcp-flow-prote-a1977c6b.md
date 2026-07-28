---
title: Create TCP Flow Protection rule.
page_id: operation-post-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-d9ecfec1
path: operations/dos-flowtrackd-api-other
description: Create a TCP Flow Protection rule for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules
operation_ids:
    - createTcpFlowProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create TCP Flow Protection rule.

`POST /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules`

Operation ID: `createTcpFlowProtectionRule`

Create a TCP Flow Protection rule for an account.

## Definition

```yaml
{"operationId": "createTcpFlowProtectionRule", "summary": "Create TCP Flow Protection rule.", "description": "Create a TCP Flow Protection rule for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The new TCP Flow Protection rule.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_NewTcpFlowProtectionRule"}}}}, "responses": {"200": {"description": "Create TCP Flow Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_tcp-flow-protection-rule-response"}}}}, "4XX": {"description": "Create TCP Flow Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
