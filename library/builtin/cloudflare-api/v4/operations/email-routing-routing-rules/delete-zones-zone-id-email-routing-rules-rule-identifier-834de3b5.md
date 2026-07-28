---
title: Delete routing rule
page_id: operation-delete-zones-zone-id-email-routing-rules-rule-identifier-cede776b
path: operations/email-routing-routing-rules
description: Delete a specific routing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/email/routing/rules/{rule_identifier}
operation_ids:
    - email-routing-routing-rules-delete-routing-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete routing rule

`DELETE /zones/{zone_id}/email/routing/rules/{rule_identifier}`

Operation ID: `email-routing-routing-rules-delete-routing-rule`

Delete a specific routing rule.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-delete-routing-rule", "summary": "Delete routing rule", "description": "Delete a specific routing rule.", "parameters": [{"name": "rule_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_rule_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Delete routing rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_rule_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
