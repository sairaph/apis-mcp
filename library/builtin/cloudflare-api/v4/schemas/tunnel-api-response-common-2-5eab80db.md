---
title: tunnel_api-response-common-2
page_id: schema-tunnel-api-response-common-2-5eab80db
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_api-response-common-2

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/tunnel_messages-2"}, "messages": {"$ref": "#/components/schemas/tunnel_messages-2"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
