---
title: mcn_response_collection
page_id: schema-mcn-response-collection-30e6f293
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_response_collection

```yaml
{"type": "object", "properties": {"messages": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_error"}}, "result_info": {"$ref": "#/components/schemas/mcn_result_info"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}
```
