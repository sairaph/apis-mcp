---
title: tls-certificates-and-hostnames_hostname_assoc_response_collection
page_id: schema-tls-certificates-and-hostnames-hostname-assoc-response-collection-52ea1184
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_hostname_assoc_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_assoc_object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Total results returned based on your search parameters.", "type": "number", "example": 1}, "page": {"description": "Current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Number of results per page of results.", "type": "number", "example": 50}, "total_count": {"description": "Total results available without any search parameters.", "type": "number", "example": 1}, "total_pages": {"description": "Total pages available of results.", "type": "number", "example": 1}}}}, "type": "object"}]}
```
