---
title: Disable Email Routing
page_id: operation-delete-zones-zone-id-email-routing-dns-bba0d9c1
path: operations/email-routing-settings
description: Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/email/routing/dns
operation_ids:
    - email-routing-settings-disable-email-routing-dns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable Email Routing

`DELETE /zones/{zone_id}/email/routing/dns`

Operation ID: `email-routing-settings-disable-email-routing-dns`

Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.

## Definition

```yaml
{"operationId": "email-routing-settings-disable-email-routing-dns", "summary": "Disable Email Routing", "description": "Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_email_setting_dns_request_body"}}}}, "responses": {"200": {"description": "Disable Email Routing response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/email_api-response-single"}, {"$ref": "#/components/schemas/email_dns_settings_response_collection"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.dns", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
