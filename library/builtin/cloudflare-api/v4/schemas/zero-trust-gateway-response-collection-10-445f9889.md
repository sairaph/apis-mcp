---
title: zero-trust-gateway_response_collection-10
page_id: schema-zero-trust-gateway-response-collection-10-445f9889
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_response_collection-10

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-4"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-7"}, "slug": {"$ref": "#/components/schemas/zero-trust-gateway_slug"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "url": {"$ref": "#/components/schemas/zero-trust-gateway_url"}}, "type": "object"}}}, "type": "object"}]}
```
