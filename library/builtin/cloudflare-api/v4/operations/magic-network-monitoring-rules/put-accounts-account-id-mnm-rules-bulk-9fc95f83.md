---
title: Update rules in bulk
page_id: operation-put-accounts-account-id-mnm-rules-bulk-1331d7e1
path: operations/magic-network-monitoring-rules
description: Update multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/mnm/rules/bulk
operation_ids:
    - magic-network-monitoring-rules-update-rules-bulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update rules in bulk

`PUT /accounts/{account_id}/mnm/rules/bulk`

Operation ID: `magic-network-monitoring-rules-update-rules-bulk`

Update multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-update-rules-bulk", "summary": "Update rules in bulk", "description": "Update multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_create"}, "maxItems": 100, "minItems": 1}}}}, "responses": {"200": {"description": "Update rules in bulk response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}}}}, "4XX": {"description": "Update rules in bulk response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write"]}
```
