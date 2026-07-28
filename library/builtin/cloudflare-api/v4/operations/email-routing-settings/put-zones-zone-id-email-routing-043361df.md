---
title: Update Email Routing settings
page_id: operation-put-zones-zone-id-email-routing-c3f3f461
path: operations/email-routing-settings
description: Update the settings for your Email Routing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/email/routing
operation_ids:
    - email-routing-settings-replace-email-routing-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Email Routing settings

`PUT /zones/{zone_id}/email/routing`

Operation ID: `email-routing-settings-replace-email-routing-settings`

Update the settings for your Email Routing zone.

## Definition

```yaml
{"operationId": "email-routing-settings-replace-email-routing-settings", "summary": "Update Email Routing settings", "description": "Update the settings for your Email Routing zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_update_email_routing_settings_properties"}}}}, "responses": {"200": {"description": "Update Email Routing settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
