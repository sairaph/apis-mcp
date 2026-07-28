---
title: Get a sending subdomain
page_id: operation-get-zones-zone-id-email-sending-subdomains-subdomain-id-f5cd88b8
path: operations/email-sending-subdomains
description: Gets information for a specific sending subdomain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}
operation_ids:
    - email-sending-subdomains-get-sending-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a sending subdomain

`GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}`

Operation ID: `email-sending-subdomains-get-sending-subdomain`

Gets information for a specific sending subdomain.

## Definition

```yaml
{"operationId": "email-sending-subdomains-get-sending-subdomain", "summary": "Get a sending subdomain", "description": "Gets information for a specific sending subdomain.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get a sending subdomain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
