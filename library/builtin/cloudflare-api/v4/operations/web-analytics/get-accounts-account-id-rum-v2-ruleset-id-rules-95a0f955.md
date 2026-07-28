---
title: List rules in Web Analytics ruleset
page_id: operation-get-accounts-account-id-rum-v2-ruleset-id-rules-6167db1d
path: operations/web-analytics
description: Lists all the rules in a Web Analytics ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/v2/{ruleset_id}/rules
operation_ids:
    - web-analytics-list-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List rules in Web Analytics ruleset

`GET /accounts/{account_id}/rum/v2/{ruleset_id}/rules`

Operation ID: `web-analytics-list-rules`

Lists all the rules in a Web Analytics ruleset.

## Definition

```yaml
{"operationId": "web-analytics-list-rules", "summary": "List rules in Web Analytics ruleset", "description": "Lists all the rules in a Web Analytics ruleset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_ruleset_identifier"}}], "responses": {"200": {"description": "List of Web Analytics rules in the ruleset.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rules-response-collection"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
