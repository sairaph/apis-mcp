---
title: magic_modified_tunnels_collection_response
page_id: schema-magic-modified-tunnels-collection-response-b9f15174
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_modified_tunnels_collection_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"properties": {"modified": {"type": "boolean", "example": true}, "modified_gre_tunnels": {"type": "array", "items": {"$ref": "#/components/schemas/magic_gre-tunnel"}}}}}}]}
```
