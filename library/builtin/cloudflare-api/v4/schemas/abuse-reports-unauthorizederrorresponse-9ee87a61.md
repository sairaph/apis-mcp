---
title: abuse-reports_UnauthorizedErrorResponse
page_id: schema-abuse-reports-unauthorizederrorresponse-9ee87a61
path: schemas
description: For HTTP 401 responses, use one of these error shapes. Cloudflare's API layer rejects requests with missing or invalid credentials before they reach the Abuse Reports API and returns the standard Cloudflare API error envelope. The Abuse Reports API returns its authorization error shape when valid credentials fail the Abuse Reports permission check or the submission entitlement check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_UnauthorizedErrorResponse

For HTTP 401 responses, use one of these error shapes. Cloudflare's API layer rejects requests with missing or invalid credentials before they reach the Abuse Reports API and returns the standard Cloudflare API error envelope. The Abuse Reports API returns its authorization error shape when valid credentials fail the Abuse Reports permission check or the submission entitlement check.

```yaml
{"description": "For HTTP 401 responses, use one of these error shapes. Cloudflare's API layer rejects requests with missing or invalid credentials before they reach the Abuse Reports API and returns the standard Cloudflare API error envelope. The Abuse Reports API returns its authorization error shape when valid credentials fail the Abuse Reports permission check or the submission entitlement check.", "oneOf": [{"$ref": "#/components/schemas/abuse-reports_CloudflareAPIErrorResponse"}, {"$ref": "#/components/schemas/abuse-reports_AuthErrorResponse"}]}
```
