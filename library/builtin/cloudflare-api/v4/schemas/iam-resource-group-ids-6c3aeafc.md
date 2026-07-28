---
title: iam_resource_group_ids
page_id: schema-iam-resource-group-ids-6c3aeafc
path: schemas
description: A set of resource groups that are specified to the policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_resource_group_ids

A set of resource groups that are specified to the policy.

```yaml
{"description": "A set of resource groups that are specified to the policy.", "type": "array", "items": {"description": "A group of scoped resources.", "properties": {"id": {"$ref": "#/components/schemas/iam_resource_group_identifier"}}, "required": ["id"], "type": "object"}, "example": [{"id": "6d7f2f5f5b1d4a0e9081fdc98d432fd1"}], "title": "Resource Group IDs"}
```
