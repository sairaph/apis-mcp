---
title: iam_user_group_policy_write_body
page_id: schema-iam-user-group-policy-write-body-016e7f71
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_user_group_policy_write_body

```yaml
{"properties": {"access": {"$ref": "#/components/schemas/iam_access"}, "permission_groups": {"$ref": "#/components/schemas/iam_permission_group_ids"}, "resource_groups": {"$ref": "#/components/schemas/iam_resource_group_ids"}}, "required": ["access", "permission_groups", "resource_groups"]}
```
