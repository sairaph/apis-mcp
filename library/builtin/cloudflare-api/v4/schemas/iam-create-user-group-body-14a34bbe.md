---
title: iam_create_user_group_body
page_id: schema-iam-create-user-group-body-14a34bbe
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create_user_group_body

```yaml
{"type": "object", "properties": {"name": {"description": "Name of the User group.", "type": "string", "example": "My New User Group", "x-auditable": true}, "policies": {"description": "Policies attached to the User group", "type": "array", "items": {"$ref": "#/components/schemas/iam_user_group_policy_write_body"}, "title": "User Group Policies"}}, "required": ["name"], "title": "Create User group with a set of policies"}
```
