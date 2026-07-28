---
title: iam_permission_group_ids
page_id: schema-iam-permission-group-ids-0c9a5c0b
path: schemas
description: A set of permission groups that are specified to the policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_permission_group_ids

A set of permission groups that are specified to the policy.

```yaml
{"description": "A set of permission groups that are specified to the policy.", "type": "array", "items": {"description": "A named group of permissions that map to a group of operations against resources.", "properties": {"id": {"$ref": "#/components/schemas/iam_permission_group_identifier"}}, "required": ["id"], "type": "object"}, "example": [{"id": "c8fed203ed3043cba015a93ad1616f1f"}, {"id": "82e64a83756745bbbb1c9c2701bf816b"}], "title": "Permission Group IDs"}
```
