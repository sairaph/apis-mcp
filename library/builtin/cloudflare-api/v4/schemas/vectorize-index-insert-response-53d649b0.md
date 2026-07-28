---
title: vectorize_index-insert-response
page_id: schema-vectorize-index-insert-response-53d649b0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-insert-response

```yaml
{"type": "object", "properties": {"count": {"description": "Specifies the count of the vectors successfully inserted.", "type": "integer", "example": 768, "x-auditable": true}, "ids": {"description": "Array of vector identifiers of the vectors successfully inserted.", "type": "array", "items": {"$ref": "#/components/schemas/vectorize_vector-identifier"}}}}
```
