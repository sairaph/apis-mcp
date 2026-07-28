---
title: vectorize_index-info-response
page_id: schema-vectorize-index-info-response-8d971c91
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_index-info-response

```yaml
{"type": "object", "properties": {"dimensions": {"$ref": "#/components/schemas/vectorize_index-dimensions"}, "processedUpToDatetime": {"description": "Specifies the timestamp the last mutation batch was processed as an ISO8601 string.", "type": "string", "format": "date-time", "example": "2024-07-22T18:25:44.442097Z", "nullable": true, "readOnly": true, "x-auditable": true}, "processedUpToMutation": {"$ref": "#/components/schemas/vectorize_mutation-uuid"}, "vectorCount": {"description": "Specifies the number of vectors present in the index", "type": "integer", "example": 300000, "x-auditable": true}}}
```
