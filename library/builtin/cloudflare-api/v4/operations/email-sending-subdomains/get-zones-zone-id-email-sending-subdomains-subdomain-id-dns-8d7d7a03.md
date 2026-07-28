---
title: Get sending subdomain DNS records
page_id: operation-get-zones-zone-id-email-sending-subdomains-subdomain-id-dns-41274274
path: operations/email-sending-subdomains
description: Returns the expected DNS records for a sending subdomain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns
operation_ids:
    - email-sending-subdomains-get-sending-subdomain-dns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get sending subdomain DNS records

`GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns`

Operation ID: `email-sending-subdomains-get-sending-subdomain-dns`

Returns the expected DNS records for a sending subdomain.

## Definition

```yaml
{"operationId": "email-sending-subdomains-get-sending-subdomain-dns", "summary": "Get sending subdomain DNS records", "description": "Returns the expected DNS records for a sending subdomain.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get sending subdomain DNS records response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_dns_settings_response_collection"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains.dns", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
