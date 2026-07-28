---
title: Update catch-all rule
page_id: operation-put-zones-zone-id-email-routing-rules-catch-all-639faf3b
path: operations/email-routing-routing-rules
description: Enable or disable catch-all routing rule, or change action to forward to a specific destination address. Forward actions require exactly one verified destination address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/email/routing/rules/catch_all
operation_ids:
    - email-routing-routing-rules-update-catch-all-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update catch-all rule

`PUT /zones/{zone_id}/email/routing/rules/catch_all`

Operation ID: `email-routing-routing-rules-update-catch-all-rule`

Enable or disable catch-all routing rule, or change action to forward to a specific destination address. Forward actions require exactly one verified destination address.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-update-catch-all-rule", "summary": "Update catch-all rule", "description": "Enable or disable catch-all routing rule, or change action to forward to a specific destination address. Forward actions require exactly one verified destination address.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_update_catch_all_rule_properties"}}}}, "responses": {"200": {"description": "Update catch-all rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_catch_all_rule_response_single"}}}}, "400": {"description": "Error 2054: Destination address is not verified. A forward action destination address must be verified before it can be used in a rule."}, "422": {"description": "Error 2007: Invalid Input. Forward actions must contain exactly one destination address."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules.catch-alls", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
