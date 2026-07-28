---
title: vectorize_index-get-vectors-by-id-response
page_id: schema-vectorize-index-get-vectors-by-id-response-b5b3dcb6
path: schemas
description: Array of vectors with matching ids.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-get-vectors-by-id-response

Array of vectors with matching ids.

```yaml
{"description": "Array of vectors with matching ids.", "type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/vectorize_vector-identifier"}, "metadata": {"type": "object"}, "namespace": {"type": "string", "nullable": true}, "values": {"type": "array", "items": {"type": "number"}}}, "type": "object"}, "example": [{"id": "some-vector-id", "metadata": {"another-key": "another-value", "customer-id": 442}, "values": [0.812, 0.621, 0.261]}, {"id": "other-vector-id", "metadata": {"another-key": "with-a-value", "customer-id": 2151}, "namespace": "namespaced", "values": [0.961, 0.751, 0.661]}]}
```
