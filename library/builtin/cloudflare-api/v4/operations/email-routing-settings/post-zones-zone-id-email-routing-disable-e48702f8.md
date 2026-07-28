---
title: Disable Email Routing
page_id: operation-post-zones-zone-id-email-routing-disable-7e7e5906
path: operations/email-routing-settings
description: Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/routing/disable
operation_ids:
    - email-routing-settings-disable-email-routing
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable Email Routing

`POST /zones/{zone_id}/email/routing/disable`

Operation ID: `email-routing-settings-disable-email-routing`

Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.

## Definition

```yaml
{"operationId": "email-routing-settings-disable-email-routing", "summary": "Disable Email Routing", "description": "Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Disable Email Routing response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing", "x-fern-sdk-method-name": "disable", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation turns off email routing for a zone."}
```
