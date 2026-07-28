---
title: vectorize_index-query-v2-request
page_id: schema-vectorize-index-query-v2-request-4a3a5331
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-query-v2-request

```yaml
{"type": "object", "properties": {"filter": {"description": "A metadata filter expression used to limit nearest neighbor results.", "type": "object", "example": {"has_viewed": {"$ne": true}, "streaming_platform": "netflix"}}, "returnMetadata": {"description": "Whether to return no metadata, indexed metadata or all metadata associated with the closest vectors.", "type": "string", "default": "none", "enum": ["none", "indexed", "all"]}, "returnValues": {"description": "Whether to return the values associated with the closest vectors.", "type": "boolean", "default": false}, "topK": {"description": "The number of nearest neighbors to find.", "type": "number", "example": 5, "default": 5}, "vector": {"description": "The search vector that will be used to find the nearest neighbors.", "type": "array", "items": {"type": "number"}, "example": [0.5, 0.5, 0.5]}}, "required": ["vector"]}
```
