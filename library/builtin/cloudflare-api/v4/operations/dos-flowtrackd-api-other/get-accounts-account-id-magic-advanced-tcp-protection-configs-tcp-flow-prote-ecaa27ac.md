---
title: Get TCP Flow Protection rule.
page_id: operation-get-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-p-0237af77
path: operations/dos-flowtrackd-api-other
description: Get a TCP Flow Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/{rule_id}
operation_ids:
    - getTcpFlowProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get TCP Flow Protection rule.

`GET /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/{rule_id}`

Operation ID: `getTcpFlowProtectionRule`

Get a TCP Flow Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "getTcpFlowProtectionRule", "summary": "Get TCP Flow Protection rule.", "description": "Get a TCP Flow Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the TCP Flow Protection rule.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Get TCP Flow Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_tcp-flow-protection-rule-response"}}}}, "4XX": {"description": "Get TCP Flow Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
