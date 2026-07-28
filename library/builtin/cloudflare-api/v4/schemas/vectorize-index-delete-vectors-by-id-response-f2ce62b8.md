---
title: vectorize_index-delete-vectors-by-id-response
page_id: schema-vectorize-index-delete-vectors-by-id-response-f2ce62b8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-delete-vectors-by-id-response

```yaml
{"type": "object", "properties": {"count": {"description": "The count of the vectors successfully deleted.", "type": "integer", "example": 42}, "ids": {"description": "Array of vector identifiers of the vectors that were successfully processed for deletion.", "type": "array", "items": {"$ref": "#/components/schemas/vectorize_vector-identifier"}}}}
```
