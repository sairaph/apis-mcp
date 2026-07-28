---
title: Update rules
page_id: operation-put-accounts-account-id-mnm-rules-8176af0d
path: operations/magic-network-monitoring-rules
description: Update network monitoring rules for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/mnm/rules
operation_ids:
    - magic-network-monitoring-rules-update-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update rules

`PUT /accounts/{account_id}/mnm/rules`

Operation ID: `magic-network-monitoring-rules-update-rules`

Update network monitoring rules for account.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-update-rules", "summary": "Update rules", "description": "Update network monitoring rules for account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_create"}}}}, "responses": {"200": {"description": "Update rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}}}}, "4XX": {"description": "Update rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write"]}
```
