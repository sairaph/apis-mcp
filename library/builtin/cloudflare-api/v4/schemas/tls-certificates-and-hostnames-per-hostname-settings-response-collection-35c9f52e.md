---
title: tls-certificates-and-hostnames_per_hostname_settings_response_collection
page_id: schema-tls-certificates-and-hostnames-per-hostname-settings-response-collection-35c9f52e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_per_hostname_settings_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_created_at-4"}, "hostname": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname-3"}, "status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_status-12"}, "updated_at": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_updated_at-5"}, "value": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_value"}}, "type": "object"}}}, "type": "object"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"type": "number", "example": 1}, "page": {"type": "number", "example": 1}, "per_page": {"type": "number", "example": 50}, "total_count": {"type": "number", "example": 1}, "total_pages": {"description": "Total pages available of results.", "type": "number", "example": 1}}}}, "type": "object"}]}
```
