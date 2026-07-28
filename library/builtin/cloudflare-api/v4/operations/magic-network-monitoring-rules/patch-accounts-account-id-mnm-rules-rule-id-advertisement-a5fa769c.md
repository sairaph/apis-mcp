---
title: Update advertisement for rule
page_id: operation-patch-accounts-account-id-mnm-rules-rule-id-advertisement-6f38c249
path: operations/magic-network-monitoring-rules
description: Update advertisement for rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/mnm/rules/{rule_id}/advertisement
operation_ids:
    - magic-network-monitoring-rules-update-advertisement-for-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update advertisement for rule

`PATCH /accounts/{account_id}/mnm/rules/{rule_id}/advertisement`

Operation ID: `magic-network-monitoring-rules-update-advertisement-for-rule`

Update advertisement for rule.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-update-advertisement-for-rule", "summary": "Update advertisement for rule", "description": "Update advertisement for rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Update advertisement for rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_advertisement_single_response"}}}}, "4XX": {"description": "Update advertisement for rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_advertisement_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write"]}
```
