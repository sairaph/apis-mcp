---
title: List sending subdomains
page_id: operation-get-zones-zone-id-email-sending-subdomains-1ebd226e
path: operations/email-sending-subdomains
description: Lists all sending-enabled subdomains for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains
operation_ids:
    - email-sending-subdomains-list-sending-subdomains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List sending subdomains

`GET /zones/{zone_id}/email/sending/subdomains`

Operation ID: `email-sending-subdomains-list-sending-subdomains`

Lists all sending-enabled subdomains for the zone.

## Definition

```yaml
{"operationId": "email-sending-subdomains-list-sending-subdomains", "summary": "List sending subdomains", "description": "Lists all sending-enabled subdomains for the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "List sending subdomains response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomains_response_collection"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
