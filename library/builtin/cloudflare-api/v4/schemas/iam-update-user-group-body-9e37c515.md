---
title: iam_update_user_group_body
page_id: schema-iam-update-user-group-body-9e37c515
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_update_user_group_body

```yaml
{"type": "object", "properties": {"name": {"description": "Name of the User group.", "type": "string", "example": "My New User Group"}, "policies": {"description": "Policies attached to the User group", "type": "array", "items": {"allOf": [{"properties": {"id": {"description": "Policy identifier.", "type": "string", "example": "f267e341f3dd4697bd3b9f71dd96247f"}}, "required": ["id"], "type": "object"}, {"$ref": "#/components/schemas/iam_user_group_policy_write_body"}]}, "title": "User Group Policies"}}, "title": "Create User group with a set of policies"}
```
