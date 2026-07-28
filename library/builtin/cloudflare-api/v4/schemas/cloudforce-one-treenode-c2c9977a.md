---
title: cloudforce-one_TreeNode
page_id: schema-cloudforce-one-treenode-c2c9977a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_TreeNode

```yaml
{"type": "object", "properties": {"children": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_TreeNode"}}, "count": {"type": "number", "example": 15}, "name": {"type": "string", "example": "workers"}, "path": {"type": "string", "example": "yara/workers"}}, "required": ["name", "path", "count", "children"]}
```
