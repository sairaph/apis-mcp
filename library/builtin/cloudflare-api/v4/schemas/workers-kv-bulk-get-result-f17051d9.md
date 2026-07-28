---
title: workers-kv_bulk-get-result
page_id: schema-workers-kv-bulk-get-result-f17051d9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_bulk-get-result

```yaml
{"type": "object", "properties": {"values": {"description": "Requested keys are paired with their values in an object.", "type": "object", "example": {"key1": "value1", "key2": "value2"}, "additionalProperties": {"description": "The value associated with the key.", "oneOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}, {"additionalProperties": true, "type": "object"}]}}}}
```
