---
title: Create rules
page_id: operation-post-accounts-account-id-mnm-rules-4ff7c052
path: operations/magic-network-monitoring-rules
description: Create network monitoring rules for account. Currently only supports creating a single rule per API request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/mnm/rules
operation_ids:
    - magic-network-monitoring-rules-create-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create rules

`POST /accounts/{account_id}/mnm/rules`

Operation ID: `magic-network-monitoring-rules-create-rules`

Create network monitoring rules for account. Currently only supports creating a single rule per API request.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-create-rules", "summary": "Create rules", "description": "Create network monitoring rules for account. Currently only supports creating a single rule per API request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_create"}}}}, "responses": {"200": {"description": "Create rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}}}}, "4XX": {"description": "Create rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin"]}
```
