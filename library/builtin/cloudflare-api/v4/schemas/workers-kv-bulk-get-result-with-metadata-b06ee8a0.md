---
title: workers-kv_bulk-get-result-with-metadata
page_id: schema-workers-kv-bulk-get-result-with-metadata-b06ee8a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_bulk-get-result-with-metadata

```yaml
{"type": "object", "properties": {"values": {"description": "Requested keys are paired with their values and metadata in an object.", "type": "object", "example": {"key1": {"expiration": 1577836800, "metadata": {"someMetadataKey": "someMetadataValue"}, "value": "value1"}, "key2": {"metadata": {"anotherKey": "anotherValue"}, "value": "value2"}}, "additionalProperties": {"nullable": true, "properties": {"expiration": {"$ref": "#/components/schemas/workers-kv_expiration"}, "metadata": {"allOf": [{"$ref": "#/components/schemas/workers-kv_any"}, {"description": "The metadata associated with the key."}]}, "value": {"allOf": [{"$ref": "#/components/schemas/workers-kv_any"}, {"description": "The value associated with the key."}]}}, "required": ["value", "metadata"], "type": "object"}}}}
```
