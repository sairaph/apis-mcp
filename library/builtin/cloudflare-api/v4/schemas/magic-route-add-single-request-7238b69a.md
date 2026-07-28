---
title: magic_route_add_single_request
page_id: schema-magic-route-add-single-request-7238b69a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_route_add_single_request

```yaml
{"type": "object", "properties": {"description": {"$ref": "#/components/schemas/magic_description"}, "nexthop": {"$ref": "#/components/schemas/magic_nexthop"}, "prefix": {"$ref": "#/components/schemas/magic_prefix"}, "priority": {"$ref": "#/components/schemas/magic_priority"}, "scope": {"$ref": "#/components/schemas/magic_scope"}, "weight": {"$ref": "#/components/schemas/magic_weight"}}, "required": ["prefix", "nexthop", "priority"]}
```
