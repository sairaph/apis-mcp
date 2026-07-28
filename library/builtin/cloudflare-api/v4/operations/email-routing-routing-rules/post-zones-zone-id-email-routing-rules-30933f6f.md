---
title: Create routing rule
page_id: operation-post-zones-zone-id-email-routing-rules-20174078
path: operations/email-routing-routing-rules
description: Rules consist of a set of criteria for matching emails (such as an email being sent to a specific custom email address) plus a set of actions to take on the email (like forwarding it to a specific destination address). Forward actions require exactly one verified destination address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/routing/rules
operation_ids:
    - email-routing-routing-rules-create-routing-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create routing rule

`POST /zones/{zone_id}/email/routing/rules`

Operation ID: `email-routing-routing-rules-create-routing-rule`

Rules consist of a set of criteria for matching emails (such as an email being sent to a specific custom email address) plus a set of actions to take on the email (like forwarding it to a specific destination address). Forward actions require exactly one verified destination address.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-create-routing-rule", "summary": "Create routing rule", "description": "Rules consist of a set of criteria for matching emails (such as an email being sent to a specific custom email address) plus a set of actions to take on the email (like forwarding it to a specific destination address). Forward actions require exactly one verified destination address.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_create_rule_properties"}}}}, "responses": {"200": {"description": "Create routing rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_rule_response_single"}}}}, "400": {"description": "Error 2054: Destination address is not verified. A forward action destination address must be verified before it can be used in a rule."}, "422": {"description": "Error 2007: Invalid Input. Forward actions must contain exactly one destination address."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
