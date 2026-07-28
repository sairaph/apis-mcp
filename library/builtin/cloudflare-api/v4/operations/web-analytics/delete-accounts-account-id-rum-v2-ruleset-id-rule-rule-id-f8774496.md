---
title: Delete a Web Analytics rule
page_id: operation-delete-accounts-account-id-rum-v2-ruleset-id-rule-rule-id-7b8608b8
path: operations/web-analytics
description: Deletes an existing rule from a Web Analytics ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rum/v2/{ruleset_id}/rule/{rule_id}
operation_ids:
    - web-analytics-delete-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Web Analytics rule

`DELETE /accounts/{account_id}/rum/v2/{ruleset_id}/rule/{rule_id}`

Operation ID: `web-analytics-delete-rule`

Deletes an existing rule from a Web Analytics ruleset.

## Definition

```yaml
{"operationId": "web-analytics-delete-rule", "summary": "Delete a Web Analytics rule", "description": "Deletes an existing rule from a Web Analytics ruleset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_ruleset_identifier"}}, {"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_rule_identifier"}}], "responses": {"200": {"description": "Deleted Web Analytics rule identifier.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rule-id-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
