---
title: api-shield_object-with-operation-id
page_id: schema-api-shield-object-with-operation-id-7d5eb792
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_object-with-operation-id

```yaml
{"type": "object", "properties": {"operation_id": {"allOf": [{"description": "The ID of the operation", "type": "string", "x-auditable": true}, {"$ref": "#/components/schemas/api-shield_uuid-2"}]}}, "required": ["operation_id"]}
```
