---
title: Create a sending subdomain
page_id: operation-post-zones-zone-id-email-sending-subdomains-3c4ada94
path: operations/email-sending-subdomains
description: Creates a new sending subdomain or re-enables sending on an existing subdomain that had it disabled. If zone-level Email Sending has not been enabled yet, the zone flag is automatically set when the entitlement is present.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains
operation_ids:
    - email-sending-subdomains-create-sending-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a sending subdomain

`POST /zones/{zone_id}/email/sending/subdomains`

Operation ID: `email-sending-subdomains-create-sending-subdomain`

Creates a new sending subdomain or re-enables sending on an existing subdomain that had it disabled. If zone-level Email Sending has not been enabled yet, the zone flag is automatically set when the entitlement is present.

## Definition

```yaml
{"operationId": "email-sending-subdomains-create-sending-subdomain", "summary": "Create a sending subdomain", "description": "Creates a new sending subdomain or re-enables sending on an existing subdomain that had it disabled. If zone-level Email Sending has not been enabled yet, the zone flag is automatically set when the entitlement is present.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_create_sending_subdomain_properties"}}}}, "responses": {"200": {"description": "Create a sending subdomain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
