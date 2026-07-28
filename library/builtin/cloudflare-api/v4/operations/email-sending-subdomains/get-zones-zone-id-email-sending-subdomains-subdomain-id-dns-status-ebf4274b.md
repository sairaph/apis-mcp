---
title: Get sending subdomain DNS status
page_id: operation-get-zones-zone-id-email-sending-subdomains-subdomain-id-dns-status-02b3fd11
path: operations/email-sending-subdomains
description: Returns the desired DNS records for a sending subdomain along with a live diff against actual DNS state. Use this to detect missing, unlocked, foreign, or multi-record conflicts before deciding whether to call the fix endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns/status
operation_ids:
    - email-sending-subdomains-get-sending-subdomain-dns-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get sending subdomain DNS status

`GET /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns/status`

Operation ID: `email-sending-subdomains-get-sending-subdomain-dns-status`

Returns the desired DNS records for a sending subdomain along with a live diff against actual DNS state. Use this to detect missing, unlocked, foreign, or multi-record conflicts before deciding whether to call the fix endpoint.

## Definition

```yaml
{"operationId": "email-sending-subdomains-get-sending-subdomain-dns-status", "summary": "Get sending subdomain DNS status", "description": "Returns the desired DNS records for a sending subdomain along with a live diff against actual DNS state. Use this to detect missing, unlocked, foreign, or multi-record conflicts before deciding whether to call the fix endpoint.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get sending subdomain DNS status response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_dns_status_response"}}}}, "400": {"description": "Error 2028: subdomain_id is not a valid UUID."}, "404": {"description": "Error 2033: sending subdomain not found."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains.dns", "x-fern-sdk-method-name": "get-status", "x-forge-hidden": true}
```
