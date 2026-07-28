---
title: Preview sending subdomain DNS
page_id: operation-post-zones-zone-id-email-sending-subdomains-preview-8fbb7395
path: operations/email-sending-subdomains
description: Returns the DNS records that would be created for a sending subdomain, flags which records are missing, and reports any conflicts with existing DNS records. This is a read-only dry-run — no records are created or modified. Use before or after creating a subdomain to check DNS status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/preview
operation_ids:
    - email-sending-subdomains-preview-sending-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview sending subdomain DNS

`POST /zones/{zone_id}/email/sending/subdomains/preview`

Operation ID: `email-sending-subdomains-preview-sending-subdomain`

Returns the DNS records that would be created for a sending subdomain, flags which records are missing, and reports any conflicts with existing DNS records. This is a read-only dry-run — no records are created or modified. Use before or after creating a subdomain to check DNS status.

## Definition

```yaml
{"operationId": "email-sending-subdomains-preview-sending-subdomain", "summary": "Preview sending subdomain DNS", "description": "Returns the DNS records that would be created for a sending subdomain, flags which records are missing, and reports any conflicts with existing DNS records. This is a read-only dry-run — no records are created or modified. Use before or after creating a subdomain to check DNS status.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_create_sending_subdomain_properties"}}}}, "responses": {"200": {"description": "Preview sending subdomain DNS response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_preview_response"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "preview", "x-forge-hidden": true}
```
