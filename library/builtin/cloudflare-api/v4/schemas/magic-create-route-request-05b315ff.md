---
title: magic_create_route_request
page_id: schema-magic-create-route-request-05b315ff
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_create_route_request

```yaml
{"type": "object", "properties": {"description": {"$ref": "#/components/schemas/magic_description"}, "nexthop": {"$ref": "#/components/schemas/magic_nexthop"}, "prefix": {"$ref": "#/components/schemas/magic_prefix"}, "priority": {"$ref": "#/components/schemas/magic_priority"}, "scope": {"$ref": "#/components/schemas/magic_scope"}, "weight": {"$ref": "#/components/schemas/magic_weight"}}, "required": ["prefix", "nexthop", "priority"]}
```
