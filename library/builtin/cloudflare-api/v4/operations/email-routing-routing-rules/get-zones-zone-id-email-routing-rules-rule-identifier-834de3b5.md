---
title: Get routing rule
page_id: operation-get-zones-zone-id-email-routing-rules-rule-identifier-97179299
path: operations/email-routing-routing-rules
description: Get information for a specific routing rule already created.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/routing/rules/{rule_identifier}
operation_ids:
    - email-routing-routing-rules-get-routing-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get routing rule

`GET /zones/{zone_id}/email/routing/rules/{rule_identifier}`

Operation ID: `email-routing-routing-rules-get-routing-rule`

Get information for a specific routing rule already created.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-get-routing-rule", "summary": "Get routing rule", "description": "Get information for a specific routing rule already created.", "parameters": [{"name": "rule_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_rule_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get routing rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_rule_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write", "Email Routing Rules Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
