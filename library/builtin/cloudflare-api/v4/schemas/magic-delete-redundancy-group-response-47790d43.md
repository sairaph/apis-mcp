---
title: magic_delete_redundancy_group_response
page_id: schema-magic-delete-redundancy-group-response-47790d43
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_delete_redundancy_group_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deleted": {"type": "boolean"}, "deleted_redundancy_group": {"$ref": "#/components/schemas/magic_redundancy_group"}}, "required": ["deleted", "deleted_redundancy_group"]}}}]}
```
