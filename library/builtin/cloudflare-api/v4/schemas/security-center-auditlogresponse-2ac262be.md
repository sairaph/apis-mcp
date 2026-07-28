---
title: security-center_auditLogResponse
page_id: schema-security-center-auditlogresponse-2ac262be
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_auditLogResponse

```yaml
{"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/security-center_auditLog"}}, "result_info": {"type": "object", "properties": {"count": {"description": "The number of items in the current result set.", "type": "integer", "example": 25}, "cursor": {"description": "Opaque cursor for the next page of results. Absent when there are no more results.", "type": "string"}, "per_page": {"description": "The requested number of items per page.", "type": "integer", "example": 25}}}}, "type": "object"}]}
```
