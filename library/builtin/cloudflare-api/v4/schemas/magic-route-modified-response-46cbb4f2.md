---
title: magic_route_modified_response
page_id: schema-magic-route-modified-response-46cbb4f2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_route_modified_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"modified": {"type": "boolean", "example": true}, "modified_route": {"$ref": "#/components/schemas/magic_route"}}}}}]}
```
