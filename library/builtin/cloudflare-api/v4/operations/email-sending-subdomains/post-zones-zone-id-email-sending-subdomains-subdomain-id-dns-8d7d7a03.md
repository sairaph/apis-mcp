---
title: Fix sending subdomain DNS records
page_id: operation-post-zones-zone-id-email-sending-subdomains-subdomain-id-dns-de90ff0b
path: operations/email-sending-subdomains
description: Idempotently re-applies the sending DNS records (creates missing records, re-applies the email_routing lock on records whose lock has been cleared). Refuses with a 409 if foreign MX, multiple SPF, multiple DMARC, or multiple DKIM records exist at the relevant DNS names — those require manual cleanup.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns
operation_ids:
    - email-sending-subdomains-fix-sending-subdomain-dns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fix sending subdomain DNS records

`POST /zones/{zone_id}/email/sending/subdomains/{subdomain_id}/dns`

Operation ID: `email-sending-subdomains-fix-sending-subdomain-dns`

Idempotently re-applies the sending DNS records (creates missing records, re-applies the email_routing lock on records whose lock has been cleared). Refuses with a 409 if foreign MX, multiple SPF, multiple DMARC, or multiple DKIM records exist at the relevant DNS names — those require manual cleanup.

## Definition

```yaml
{"operationId": "email-sending-subdomains-fix-sending-subdomain-dns", "summary": "Fix sending subdomain DNS records", "description": "Idempotently re-applies the sending DNS records (creates missing records, re-applies the email_routing lock on records whose lock has been cleared). Refuses with a 409 if foreign MX, multiple SPF, multiple DMARC, or multiple DKIM records exist at the relevant DNS names — those require manual cleanup.", "parameters": [{"name": "subdomain_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_sending_subdomain_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Fix sending subdomain DNS records response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_subdomain_dns_status_response"}}}}, "400": {"description": "Error 2028: subdomain_id is not a valid UUID."}, "403": {"description": "Error 2043: zone is admin-locked and cannot be mutated."}, "404": {"description": "Error 2033: sending subdomain not found."}, "409": {"description": "Conflict — manual DNS cleanup required. Errors 2008 (foreign MX), 2026 (multiple SPF), 2027 (multiple DMARC), or a multi-DKIM error."}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending subdomains"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.subdomains.dns", "x-fern-sdk-method-name": "fix", "x-forge-hidden": true}
```
