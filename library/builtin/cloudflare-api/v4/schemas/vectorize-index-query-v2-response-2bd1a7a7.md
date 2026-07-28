---
title: vectorize_index-query-v2-response
page_id: schema-vectorize-index-query-v2-response-2bd1a7a7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-query-v2-response

```yaml
{"type": "object", "properties": {"count": {"description": "Specifies the count of vectors returned by the search", "type": "integer"}, "matches": {"description": "Array of vectors matched by the search", "type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/vectorize_vector-identifier"}, "metadata": {"type": "object", "nullable": true}, "namespace": {"type": "string", "nullable": true}, "score": {"description": "The score of the vector according to the index's distance metric", "type": "number"}, "values": {"type": "array", "items": {"type": "number"}, "nullable": true}}, "type": "object"}}}}
```
