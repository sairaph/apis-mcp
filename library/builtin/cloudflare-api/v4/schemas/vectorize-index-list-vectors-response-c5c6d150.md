---
title: vectorize_index-list-vectors-response
page_id: schema-vectorize-index-list-vectors-response-c5c6d150
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-list-vectors-response

```yaml
{"type": "object", "properties": {"count": {"description": "Number of vectors returned in this response", "type": "integer", "example": 100}, "cursorExpirationTimestamp": {"description": "When the cursor expires as an ISO8601 string", "type": "string", "format": "date-time", "example": "2025-08-12T20:32:52.469144957+00:00", "nullable": true}, "isTruncated": {"description": "Whether there are more vectors available beyond this response", "type": "boolean", "example": true}, "nextCursor": {"description": "Cursor for the next page of results", "type": "string", "example": "suUTaDY5PFUiRweVccnzyt9n75suNPbXHPshvCzue5mHjtj7Letjvzlza9eGj099", "nullable": true}, "totalCount": {"description": "Total number of vectors in the index", "type": "integer", "example": 500}, "vectors": {"description": "Array of vector items", "type": "array", "items": {"$ref": "#/components/schemas/vectorize_vector-list-item"}}}, "required": ["count", "totalCount", "isTruncated", "vectors"]}
```
