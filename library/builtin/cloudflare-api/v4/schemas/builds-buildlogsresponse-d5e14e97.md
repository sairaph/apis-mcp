---
title: builds_BuildLogsResponse
page_id: schema-builds-buildlogsresponse-d5e14e97
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_BuildLogsResponse

```yaml
{"type": "object", "properties": {"cursor": {"$ref": "#/components/schemas/builds_cursor"}, "lines": {"type": "array", "items": {"items": {"anyOf": [{"description": "Unix epoch timestamp", "example": 1636472400, "type": "number"}, {"description": "Log message", "example": "Building worker...", "type": "string"}]}, "maxItems": 2, "minItems": 2, "type": "array"}}, "truncated": {"type": "boolean", "example": false}}}
```
