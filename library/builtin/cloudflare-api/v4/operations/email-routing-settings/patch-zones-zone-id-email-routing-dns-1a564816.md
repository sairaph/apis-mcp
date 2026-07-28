---
title: Unlock Email Routing
page_id: operation-patch-zones-zone-id-email-routing-dns-345ea2c4
path: operations/email-routing-settings
description: Unlock MX Records previously locked by Email Routing.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/email/routing/dns
operation_ids:
    - email-routing-settings-unlock-email-routing-dns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unlock Email Routing

`PATCH /zones/{zone_id}/email/routing/dns`

Operation ID: `email-routing-settings-unlock-email-routing-dns`

Unlock MX Records previously locked by Email Routing.

## Definition

```yaml
{"operationId": "email-routing-settings-unlock-email-routing-dns", "summary": "Unlock Email Routing", "description": "Unlock MX Records previously locked by Email Routing.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_setting_dns_request_body"}}}}, "responses": {"200": {"description": "Unlock Email Routing MX records", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_settings_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.dns", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
