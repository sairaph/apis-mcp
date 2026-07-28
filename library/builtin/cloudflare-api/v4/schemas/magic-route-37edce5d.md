---
title: magic_route
page_id: schema-magic-route-37edce5d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_route

```yaml
{"type": "object", "properties": {"created_on": {"$ref": "#/components/schemas/magic_created_on"}, "description": {"$ref": "#/components/schemas/magic_description"}, "id": {"$ref": "#/components/schemas/magic_identifier"}, "modified_on": {"$ref": "#/components/schemas/magic_modified_on"}, "nexthop": {"$ref": "#/components/schemas/magic_nexthop"}, "prefix": {"$ref": "#/components/schemas/magic_prefix"}, "priority": {"$ref": "#/components/schemas/magic_priority"}, "scope": {"$ref": "#/components/schemas/magic_scope"}, "weight": {"$ref": "#/components/schemas/magic_weight"}}, "required": ["id", "prefix", "nexthop", "priority"]}
```
