---
title: magic_tunnel_deleted_response
page_id: schema-magic-tunnel-deleted-response-2230402b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_tunnel_deleted_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deleted": {"type": "boolean", "example": true}, "deleted_gre_tunnel": {"$ref": "#/components/schemas/magic_gre-tunnel"}}}}}]}
```
