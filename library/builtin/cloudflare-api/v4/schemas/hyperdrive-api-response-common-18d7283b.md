---
title: hyperdrive_api-response-common
page_id: schema-hyperdrive-api-response-common-18d7283b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/hyperdrive_messages"}, "messages": {"$ref": "#/components/schemas/hyperdrive_messages"}, "result": {"type": "object"}, "success": {"description": "Return the status of the API call success.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
