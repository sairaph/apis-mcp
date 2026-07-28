---
title: mconn_bad_response
page_id: schema-mconn-bad-response-97aa876a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_bad_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mconn_response"}, {"properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_coded_message"}, "minLength": 1}, "result": {"$ref": "#/components/schemas/mconn_none"}}, "type": "object"}]}
```
