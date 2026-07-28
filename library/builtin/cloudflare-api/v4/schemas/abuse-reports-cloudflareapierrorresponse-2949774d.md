---
title: abuse-reports_CloudflareAPIErrorResponse
page_id: schema-abuse-reports-cloudflareapierrorresponse-2949774d
path: schemas
description: Cloudflare's API layer uses this standard error envelope for requests that fail authentication before they reach the Abuse Reports API. The example shows the authentication error for missing or invalid credentials.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_CloudflareAPIErrorResponse

Cloudflare's API layer uses this standard error envelope for requests that fail authentication before they reach the Abuse Reports API. The example shows the authentication error for missing or invalid credentials.

```yaml
{"description": "Cloudflare's API layer uses this standard error envelope for requests that fail authentication before they reach the Abuse Reports API. The example shows the authentication error for missing or invalid credentials.", "type": "object", "properties": {"errors": {"description": "Cloudflare API error details.", "type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_ErrorMessage"}, "example": [{"code": 10000, "message": "Authentication error"}]}, "messages": {"description": "Informational messages.", "type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}, "example": []}, "result": {"description": "Always null for this error response.", "type": "object", "nullable": true}, "success": {"description": "Whether the request succeeded.", "type": "boolean", "example": false}}, "example": {"errors": [{"code": 10000, "message": "Authentication error"}], "messages": [], "result": null, "success": false}, "required": ["success", "errors", "messages", "result"]}
```
