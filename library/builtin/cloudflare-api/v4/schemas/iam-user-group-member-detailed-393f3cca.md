---
title: iam_user_group_member_detailed
page_id: schema-iam-user-group-member-detailed-393f3cca
path: schemas
description: Detailed member information for a User Group member.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_user_group_member_detailed

Detailed member information for a User Group member.

```yaml
{"description": "Detailed member information for a User Group member.", "type": "object", "properties": {"created_at": {"description": "When the member was added to the user group.", "type": "string", "format": "date-time", "example": "2026-01-15T10:30:00Z", "readOnly": true}, "email": {"$ref": "#/components/schemas/iam_email"}, "id": {"description": "Account member identifier.", "type": "string", "example": "4f5f0c14a2a41d5063dd301b2f829f04", "readOnly": true}, "status": {"description": "The member's status in the account.", "example": "accepted", "enum": ["accepted", "pending"], "readOnly": true}, "user": {"description": "Details of the user associated with this membership.", "type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "first_name": {"description": "User's first name.", "type": "string", "example": "Alice", "readOnly": true}, "id": {"description": "User identifier tag.", "type": "string", "example": "7c5dae5552338874e5053f2534d2767a", "readOnly": true}, "last_name": {"description": "User's last name.", "type": "string", "example": "Smith", "readOnly": true}}, "readOnly": true}}, "required": ["id"]}
```
