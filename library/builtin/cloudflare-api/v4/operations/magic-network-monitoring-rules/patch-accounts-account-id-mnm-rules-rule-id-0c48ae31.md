---
title: Update rule
page_id: operation-patch-accounts-account-id-mnm-rules-rule-id-fadb24c6
path: operations/magic-network-monitoring-rules
description: Update a network monitoring rule for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/mnm/rules/{rule_id}
operation_ids:
    - magic-network-monitoring-rules-update-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update rule

`PATCH /accounts/{account_id}/mnm/rules/{rule_id}`

Operation ID: `magic-network-monitoring-rules-update-rule`

Update a network monitoring rule for account.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-update-rule", "summary": "Update rule", "description": "Update a network monitoring rule for account.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_create"}}}}, "responses": {"200": {"description": "Update rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}}}}, "4XX": {"description": "Update rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write"]}
```
