---
title: Enable Email Routing
page_id: operation-post-zones-zone-id-email-routing-enable-f0b1cc56
path: operations/email-routing-settings
description: Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/routing/enable
operation_ids:
    - email-routing-settings-enable-email-routing
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable Email Routing

`POST /zones/{zone_id}/email/routing/enable`

Operation ID: `email-routing-settings-enable-email-routing`

Enable you Email Routing zone. Add and lock the necessary MX and SPF records.

## Definition

```yaml
{"operationId": "email-routing-settings-enable-email-routing", "summary": "Enable Email Routing", "description": "Enable you Email Routing zone. Add and lock the necessary MX and SPF records.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Enable Email Routing response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing", "x-fern-sdk-method-name": "enable", "x-forge-hidden": true}
```
