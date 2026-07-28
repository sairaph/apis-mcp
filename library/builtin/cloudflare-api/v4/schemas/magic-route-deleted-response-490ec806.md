---
title: magic_route_deleted_response
page_id: schema-magic-route-deleted-response-490ec806
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_route_deleted_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deleted": {"type": "boolean", "example": true}, "deleted_route": {"$ref": "#/components/schemas/magic_route"}}}}}]}
```
