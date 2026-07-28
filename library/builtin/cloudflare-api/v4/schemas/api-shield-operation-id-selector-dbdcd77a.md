---
title: api-shield_operation_id_selector
page_id: schema-api-shield-operation-id-selector-dbdcd77a
path: schemas
description: Operation IDs selector
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_id_selector

Operation IDs selector

```yaml
{"description": "Operation IDs selector", "type": "object", "properties": {"include": {"type": "object", "properties": {"operation_ids": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "minItems": 1}}, "required": ["operation_ids"]}}, "required": ["include"]}
```
