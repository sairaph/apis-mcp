---
title: List all TCP Flow Protection rules.
page_id: operation-get-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-p-deadcd47
path: operations/dos-flowtrackd-api-other
description: List all TCP Flow Protection rules for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules
operation_ids:
    - listTcpFlowProtectionRulesForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all TCP Flow Protection rules.

`GET /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules`

Operation ID: `listTcpFlowProtectionRulesForAccount`

List all TCP Flow Protection rules for an account.

## Definition

```yaml
{"operationId": "listTcpFlowProtectionRulesForAccount", "summary": "List all TCP Flow Protection rules.", "description": "List all TCP Flow Protection rules for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "page", "in": "query", "description": "The page number for pagination. Defaults to 1.", "schema": {"type": "integer", "format": "int64", "x-auditable": true}, "explode": false}, {"name": "per_page", "in": "query", "description": "The number of items per page. Must be between 10 and 1000. Defaults to 25.", "schema": {"type": "integer", "format": "int64", "x-auditable": true}, "explode": false}, {"name": "order", "in": "query", "description": "The field to order by. Defaults to 'prefix'.", "schema": {"type": "string", "x-auditable": true}, "explode": false}, {"name": "direction", "in": "query", "description": "The direction of ordering (ASC or DESC). Defaults to 'ASC'.", "schema": {"type": "string", "x-auditable": true}, "explode": false}], "responses": {"200": {"description": "List all TCP Flow Protection rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_tcp-flow-protection-rule-list-response"}}}}, "4XX": {"description": "List all TCP Flow Protection rules failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
