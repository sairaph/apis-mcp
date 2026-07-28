---
title: speed_api-response-single-id
page_id: schema-speed-api-response-single-id-f7372dd1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# speed_api-response-single-id

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/speed_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true, "properties": {"id": {"$ref": "#/components/schemas/speed_identifier"}}, "required": ["id"]}}, "type": "object"}]}
```
