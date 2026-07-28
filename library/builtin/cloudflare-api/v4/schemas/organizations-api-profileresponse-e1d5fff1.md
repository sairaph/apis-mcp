---
title: organizations-api_ProfileResponse
page_id: schema-organizations-api-profileresponse-e1d5fff1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_ProfileResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Profile"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
