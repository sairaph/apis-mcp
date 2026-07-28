---
title: List routing rules
page_id: operation-get-zones-zone-id-email-routing-rules-0b1f1338
path: operations/email-routing-routing-rules
description: Lists existing routing rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/routing/rules
operation_ids:
    - email-routing-routing-rules-list-routing-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List routing rules

`GET /zones/{zone_id}/email/routing/rules`

Operation ID: `email-routing-routing-rules-list-routing-rules`

Lists existing routing rules.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-list-routing-rules", "summary": "List routing rules", "description": "Lists existing routing rules.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "enabled", "in": "query", "schema": {"description": "Filter by enabled routing rules.", "type": "boolean", "example": true, "enum": [true, false]}}], "responses": {"200": {"description": "List routing rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_rules_response_collection"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": ["Email Routing Rules Write", "Email Routing Rules Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.rule.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
