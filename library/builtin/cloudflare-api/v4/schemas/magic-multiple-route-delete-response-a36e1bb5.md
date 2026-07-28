---
title: magic_multiple_route_delete_response
page_id: schema-magic-multiple-route-delete-response-a36e1bb5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_multiple_route_delete_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"properties": {"deleted": {"type": "boolean", "example": true}, "deleted_routes": {"type": "array", "items": {"$ref": "#/components/schemas/magic_route"}}}}}}]}
```
