---
title: magic_redundancy_group_with_members
page_id: schema-magic-redundancy-group-with-members-b7992b52
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_redundancy_group_with_members

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/magic_redundancy_group"}, {"properties": {"member_data": {"type": "array", "items": {"$ref": "#/components/schemas/magic_redundancy_member_data"}}}}]}
```
