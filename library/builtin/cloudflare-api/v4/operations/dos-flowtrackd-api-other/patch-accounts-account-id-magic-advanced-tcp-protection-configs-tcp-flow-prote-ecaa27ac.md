---
title: Update TCP Flow Protection rule.
page_id: operation-patch-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-fc5f26b1
path: operations/dos-flowtrackd-api-other
description: Update a TCP Flow Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/{rule_id}
operation_ids:
    - updateTcpFlowProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update TCP Flow Protection rule.

`PATCH /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/{rule_id}`

Operation ID: `updateTcpFlowProtectionRule`

Update a TCP Flow Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "updateTcpFlowProtectionRule", "summary": "Update TCP Flow Protection rule.", "description": "Update a TCP Flow Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the TCP Flow Protection rule to update.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "requestBody": {"description": "The updates to apply to the TCP Flow Protection rule.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_TcpFlowProtectionRuleUpdate"}}}}, "responses": {"200": {"description": "Update TCP Flow Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_tcp-flow-protection-rule-response"}}}}, "4XX": {"description": "Update TCP Flow Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
