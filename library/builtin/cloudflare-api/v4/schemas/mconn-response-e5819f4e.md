---
title: mconn_response
page_id: schema-mconn-response-e5819f4e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_response

```yaml
{"type": "object", "properties": {"messages": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_coded_message"}}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}
```
