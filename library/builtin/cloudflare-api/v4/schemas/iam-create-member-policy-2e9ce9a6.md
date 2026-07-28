---
title: iam_create_member_policy
page_id: schema-iam-create-member-policy-2e9ce9a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create_member_policy

```yaml
{"properties": {"access": {"$ref": "#/components/schemas/iam_access"}, "id": {"$ref": "#/components/schemas/iam_policy_identifier"}, "permission_groups": {"$ref": "#/components/schemas/iam_member_permission_groups"}, "resource_groups": {"$ref": "#/components/schemas/iam_member_resource_groups"}}, "required": ["id", "access", "permission_groups", "resource_groups"], "title": "create_member_policy"}
```
