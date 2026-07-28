---
title: vectorize_index-get-vectors-by-id-request
page_id: schema-vectorize-index-get-vectors-by-id-request-f43380a3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-get-vectors-by-id-request

```yaml
{"type": "object", "properties": {"ids": {"description": "A list of vector identifiers to retrieve from the index indicated by the path.", "type": "array", "items": {"$ref": "#/components/schemas/vectorize_vector-identifier"}, "example": ["5121db81354a40c6aedc3fe1ace51c59", "f90eb49c2107486abdfd78c67e853430"]}}}
```
