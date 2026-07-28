---
title: List rules
page_id: operation-get-accounts-account-id-mnm-rules-c4fca936
path: operations/magic-network-monitoring-rules
description: Lists network monitoring rules for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mnm/rules
operation_ids:
    - magic-network-monitoring-rules-list-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List rules

`GET /accounts/{account_id}/mnm/rules`

Operation ID: `magic-network-monitoring-rules-list-rules`

Lists network monitoring rules for account.

## Definition

```yaml
{"operationId": "magic-network-monitoring-rules-list-rules", "summary": "List rules", "description": "Lists network monitoring rules for account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "responses": {"200": {"description": "List rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}}}}, "4XX": {"description": "List rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rules_collection_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Rules"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write", "Magic Network Monitoring Config Read"]}
```
