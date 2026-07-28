---
title: Create a Web Analytics rule
page_id: operation-post-accounts-account-id-rum-v2-ruleset-id-rule-48276b6b
path: operations/web-analytics
description: Creates a new rule in a Web Analytics ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rum/v2/{ruleset_id}/rule
operation_ids:
    - web-analytics-create-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Web Analytics rule

`POST /accounts/{account_id}/rum/v2/{ruleset_id}/rule`

Operation ID: `web-analytics-create-rule`

Creates a new rule in a Web Analytics ruleset.

## Definition

```yaml
{"operationId": "web-analytics-create-rule", "summary": "Create a Web Analytics rule", "description": "Creates a new rule in a Web Analytics ruleset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_ruleset_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_create-rule-request"}}}}, "responses": {"200": {"description": "Created Web Analytics rule.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rule-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
