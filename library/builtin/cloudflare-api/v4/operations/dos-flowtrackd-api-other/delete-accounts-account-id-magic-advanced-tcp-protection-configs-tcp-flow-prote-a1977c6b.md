---
title: Delete all TCP Flow Protection rules.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flo-85937b75
path: operations/dos-flowtrackd-api-other
description: Delete all TCP Flow Protection rules for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules
operation_ids:
    - deleteTcpFlowProtectionRulesForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all TCP Flow Protection rules.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules`

Operation ID: `deleteTcpFlowProtectionRulesForAccount`

Delete all TCP Flow Protection rules for an account.

## Definition

```yaml
{"operationId": "deleteTcpFlowProtectionRulesForAccount", "summary": "Delete all TCP Flow Protection rules.", "description": "Delete all TCP Flow Protection rules for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Delete all TCP Flow Protection rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete all TCP Flow Protection rules failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
