---
title: Get catch-all rule
page_id: operation-get-zones-zone-id-email-routing-rules-catch-all-6328cf33
path: operations/email-routing-routing-rules
description: Get information on the default catch-all routing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/routing/rules/catch_all
operation_ids:
    - email-routing-routing-rules-get-catch-all-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get catch-all rule

`GET /zones/{zone_id}/email/routing/rules/catch_all`

Operation ID: `email-routing-routing-rules-get-catch-all-rule`

Get information on the default catch-all routing rule.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-get-catch-all-rule", "summary": "Get catch-all rule", "description": "Get information on the default catch-all routing rule.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get catch-all rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_catch_all_rule_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write", "Email Routing Rules Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules.catch-alls", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
