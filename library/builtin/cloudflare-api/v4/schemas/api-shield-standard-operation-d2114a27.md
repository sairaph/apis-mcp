---
title: api-shield_standard_operation
page_id: schema-api-shield-standard-operation-d2114a27
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_standard_operation

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_basic_operation"}, {"properties": {"last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "operation_id": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, "required": ["operation_id", "last_updated"], "type": "object"}]}
```
