---
title: access_response_collection-24
page_id: schema-access-response-collection-24-3c1d81a2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_response_collection-24

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"example": 1}, "page": {"example": 1}, "per_page": {"example": 100}, "total_count": {"example": 1}}}}}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/access_users-2"}}}}]}
```
