---
title: mcn_bad_response
page_id: schema-mcn-bad-response-88843f52
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_bad_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mcn_response"}, {"properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_error"}, "minLength": 1}, "result": {"type": "object", "enum": [null], "nullable": true}}, "type": "object"}]}
```
