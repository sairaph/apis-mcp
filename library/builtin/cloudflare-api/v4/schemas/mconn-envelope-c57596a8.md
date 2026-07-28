---
title: mconn_envelope
page_id: schema-mconn-envelope-c57596a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_envelope

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_coded_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_coded_message"}}, "success": {"type": "boolean"}}, "required": ["success"]}
```
