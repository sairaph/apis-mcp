---
title: Email Routing - DNS settings
page_id: operation-get-zones-zone-id-email-routing-dns-e42b4571
path: operations/email-routing-settings
description: Show the DNS records needed to configure your Email Routing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/routing/dns
operation_ids:
    - email-routing-settings-email-routing-dns-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Email Routing - DNS settings

`GET /zones/{zone_id}/email/routing/dns`

Operation ID: `email-routing-settings-email-routing-dns-settings`

Show the DNS records needed to configure your Email Routing zone.

## Definition

```yaml
{"operationId": "email-routing-settings-email-routing-dns-settings", "summary": "Email Routing - DNS settings", "description": "Show the DNS records needed to configure your Email Routing zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}, {"name": "subdomain", "in": "query", "schema": {"$ref": "#/components/schemas/email_email_setting_name"}}], "responses": {"200": {"description": "Email Routing - DNS settings response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/email_email_routing_dns_query_response"}, {"$ref": "#/components/schemas/email_dns_settings_response_collection"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.zone.email.routing.config.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.dns", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
