---
title: iam_user_group_member
page_id: schema-iam-user-group-member-38a5a148
path: schemas
description: Member attached to a User Group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_user_group_member

Member attached to a User Group.

```yaml
{"description": "Member attached to a User Group.", "type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "id": {"description": "Account member identifier.", "type": "string", "example": "4f5f0c14a2a41d5063dd301b2f829f04", "readOnly": true}, "status": {"description": "The member's status in the account.", "example": "accepted", "enum": ["accepted", "pending"], "readOnly": true}}, "required": ["id"]}
```
