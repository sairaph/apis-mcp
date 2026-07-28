---
title: iam_scim_error_response
page_id: schema-iam-scim-error-response-3464c950
path: schemas
description: SCIM error envelope (RFC 7644 Section 3.12). Returned on all 4XX and 5XX responses from SCIM endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_error_response

SCIM error envelope (RFC 7644 Section 3.12). Returned on all 4XX and 5XX responses from SCIM endpoints.

```yaml
{"description": "SCIM error envelope (RFC 7644 Section 3.12). Returned on all 4XX and 5XX responses from SCIM endpoints.\n", "type": "object", "properties": {"detail": {"description": "A human-readable message describing the error.", "type": "string", "example": "Invalid SCIM request: userName is required"}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:api:messages:2.0:Error"]}, "scimType": {"description": "A SCIM detail error keyword (RFC 7644 Section 3.12). Only present for specific error conditions (e.g. `uniqueness` for duplicate member).\n", "type": "string", "example": "uniqueness"}, "status": {"description": "The HTTP status code as a string.", "type": "string", "example": "400"}}, "required": ["schemas", "status"], "title": "SCIM Error Response"}
```
