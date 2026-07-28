---
title: vectorize_list-metadata-index-response
page_id: schema-vectorize-list-metadata-index-response-036a78fc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_list-metadata-index-response

```yaml
{"type": "object", "properties": {"metadataIndexes": {"description": "Array of indexed metadata properties.", "type": "array", "items": {"properties": {"indexType": {"description": "Specifies the type of indexed metadata property.", "type": "string", "enum": ["string", "number", "boolean"], "x-auditable": true}, "propertyName": {"description": "Specifies the indexed metadata property.", "type": "string", "example": "random_metadata_property", "x-auditable": true}}, "type": "object"}}}, "example": {"metadataIndexes": [{"indexType": "number", "propertyName": "some-num-prop"}, {"indexType": "string", "propertyName": "some-str-prop"}, {"indexType": "boolean", "propertyName": "some-bool-prop"}]}}
```
