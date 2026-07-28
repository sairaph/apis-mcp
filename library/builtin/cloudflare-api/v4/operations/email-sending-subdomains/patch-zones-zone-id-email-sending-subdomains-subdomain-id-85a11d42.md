---
title: Update a sending subdomain
page_id: operation-patch-zones-zone-id-email-sending-subdomains-subdomain-id-a74513f0
path: operations/email-sending-subdomains
description: Updates the activity-log preview preference for a sending subdomain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}
operation_ids:
    - email-sending-subdomains-update-sending-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a sending subdomain

`PATCH /zones/{zone_id}/email/sending/subdomains/{subdomain_id}`

Operation ID: `email-sending-subdomains-update-sending-subdomain`

Updates the activity-log preview preference for a sending subdomain.

## Definition

```yaml
{"operationId": "email-sending-subdomains-update-sending-subdomain", "summary": "Update a sending subdomain", "description": "Updates the activity-log preview preference for a sending subdomain.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_update_sending_subdomain_properties"}}}}, "responses": {"200": {"description": "Update a sending subdomain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
