---
title: tls-certificates-and-hostnames_certificate_response_collection-5
page_id: schema-tls-certificates-and-hostnames-certificate-response-collection-5-fc4c9d4d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_certificate_response_collection-5

```yaml
{"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificateObject-3"}}}, "type": "object"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"type": "number", "example": 1}, "page": {"type": "number", "example": 1}, "per_page": {"type": "number", "example": 50}, "total_count": {"type": "number", "example": 1}, "total_pages": {"description": "Total pages available of results.", "type": "number", "example": 1}}}}, "type": "object"}]}
```
