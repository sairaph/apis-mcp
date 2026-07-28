---
title: Update Web Analytics rules
page_id: operation-post-accounts-account-id-rum-v2-ruleset-id-rules-4ef28ba8
path: operations/web-analytics
description: Modifies one or more rules in a Web Analytics ruleset with a single request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rum/v2/{ruleset_id}/rules
operation_ids:
    - web-analytics-modify-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Web Analytics rules

`POST /accounts/{account_id}/rum/v2/{ruleset_id}/rules`

Operation ID: `web-analytics-modify-rules`

Modifies one or more rules in a Web Analytics ruleset with a single request.

## Definition

```yaml
{"operationId": "web-analytics-modify-rules", "summary": "Update Web Analytics rules", "description": "Modifies one or more rules in a Web Analytics ruleset with a single request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_ruleset_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_modify-rules-request"}}}}, "responses": {"200": {"description": "List of modified Web Analytics rules.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rules-response-collection"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.rules", "x-fern-sdk-method-name": "bulk-create", "x-forge-hidden": true}
```
