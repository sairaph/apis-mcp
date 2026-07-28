---
title: Delete a sending subdomain
page_id: operation-delete-zones-zone-id-email-sending-subdomains-subdomain-id-7717e8fe
path: operations/email-sending-subdomains
description: Disables sending on a subdomain and removes its DNS records. If routing is still active on the subdomain, only sending is disabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}
operation_ids:
    - email-sending-subdomains-delete-sending-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a sending subdomain

`DELETE /zones/{zone_id}/email/sending/subdomains/{subdomain_id}`

Operation ID: `email-sending-subdomains-delete-sending-subdomain`

Disables sending on a subdomain and removes its DNS records. If routing is still active on the subdomain, only sending is disabled.

## Definition

```yaml
{"operationId": "email-sending-subdomains-delete-sending-subdomain", "summary": "Delete a sending subdomain", "description": "Disables sending on a subdomain and removes its DNS records. If routing is still active on the subdomain, only sending is disabled.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Delete a sending subdomain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_api-response-single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
