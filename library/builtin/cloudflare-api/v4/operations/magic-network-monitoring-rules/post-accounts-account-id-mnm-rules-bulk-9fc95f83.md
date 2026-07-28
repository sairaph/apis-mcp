---
title: Create rules in bulk
page_id: operation-post-accounts-account-id-mnm-rules-bulk-181dfc06
path: operations/magic-network-monitoring-rules
description: Create multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/mnm/rules/bulk
operation_ids:
    - magic-network-monitoring-rules-create-rules-bulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create rules in bulk

`POST /accounts/{account_id}/mnm/rules/bulk`

Operation ID: `magic-network-monitoring-rules-create-rules-bulk`

Create multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-create-rules-bulk", "summary": "Create rules in bulk", "description": "Create multiple network monitoring rules for account in a single request. Supports up to 100 rules per request. All rules in a single request must be of the same type.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_create"}, "maxItems": 100, "minItems": 1}}}}, "responses": {"200": {"description": "Create rules in bulk response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}}}}, "4XX": {"description": "Create rules in bulk response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin"]}
```
