---
title: Get Email Routing settings
page_id: operation-get-zones-zone-id-email-routing-1dd2e8b5
path: operations/email-routing-settings
description: Get information about the settings for your Email Routing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/routing
operation_ids:
    - email-routing-settings-get-email-routing-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Email Routing settings

`GET /zones/{zone_id}/email/routing`

Operation ID: `email-routing-settings-get-email-routing-settings`

Get information about the settings for your Email Routing zone.

## Definition

```yaml
{"operationId": "email-routing-settings-get-email-routing-settings", "summary": "Get Email Routing settings", "description": "Get information about the settings for your Email Routing zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get Email Routing settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
