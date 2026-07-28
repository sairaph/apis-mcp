---
title: Get rule
page_id: operation-get-accounts-account-id-mnm-rules-rule-id-bcfc2bbc
path: operations/magic-network-monitoring-rules
description: List a single network monitoring rule for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mnm/rules/{rule_id}
operation_ids:
    - magic-network-monitoring-rules-get-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get rule

`GET /accounts/{account_id}/mnm/rules/{rule_id}`

Operation ID: `magic-network-monitoring-rules-get-rule`

List a single network monitoring rule for account.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-get-rule", "summary": "Get rule", "description": "List a single network monitoring rule for account.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "responses": {"200": {"description": "Get rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}}}}, "4XX": {"description": "Get rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write", "Magic Network Monitoring Config Read"]}
```
