---
title: Unlock Email Routing
page_id: operation-post-zones-zone-id-email-routing-unlock-30590eca
path: operations/email-routing-settings
description: Unlock MX records previously locked by Email Routing. Deprecated - use PATCH /zones/{zone_id}/email/routing/dns instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/routing/unlock
operation_ids:
    - email-routing-settings-unlock-email-routing
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unlock Email Routing

`POST /zones/{zone_id}/email/routing/unlock`

Operation ID: `email-routing-settings-unlock-email-routing`

Unlock MX records previously locked by Email Routing. Deprecated - use PATCH /zones/{zone_id}/email/routing/dns instead.

## Definition

```yaml
{"operationId": "email-routing-settings-unlock-email-routing", "summary": "Unlock Email Routing", "description": "Unlock MX records previously locked by Email Routing. Deprecated - use PATCH /zones/{zone_id}/email/routing/dns instead.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_setting_dns_request_body"}}}}, "responses": {"200": {"description": "Unlock Email Routing response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing", "x-fern-sdk-method-name": "unlock", "x-forge-hidden": true}
```
