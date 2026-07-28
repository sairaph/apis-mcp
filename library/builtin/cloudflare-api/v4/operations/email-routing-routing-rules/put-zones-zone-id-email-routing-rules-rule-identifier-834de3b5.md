---
title: Update routing rule
page_id: operation-put-zones-zone-id-email-routing-rules-rule-identifier-f43fe722
path: operations/email-routing-routing-rules
description: Update actions and matches, or enable/disable specific routing rules. Forward actions require exactly one verified destination address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/email/routing/rules/{rule_identifier}
operation_ids:
    - email-routing-routing-rules-update-routing-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update routing rule

`PUT /zones/{zone_id}/email/routing/rules/{rule_identifier}`

Operation ID: `email-routing-routing-rules-update-routing-rule`

Update actions and matches, or enable/disable specific routing rules. Forward actions require exactly one verified destination address.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-update-routing-rule", "summary": "Update routing rule", "description": "Update actions and matches, or enable/disable specific routing rules. Forward actions require exactly one verified destination address.", "parameters": [{"name": "rule_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_rule_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_update_rule_properties"}}}}, "responses": {"200": {"description": "Update routing rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_rule_response_single"}}}}, "400": {"description": "Error 2054: Destination address is not verified. A forward action destination address must be verified before it can be used in a rule."}, "422": {"description": "Error 2007: Invalid Input. Forward actions must contain exactly one destination address."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
