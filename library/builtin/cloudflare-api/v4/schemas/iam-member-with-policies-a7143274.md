---
title: iam_member_with_policies
page_id: schema-iam-member-with-policies-a7143274
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_member_with_policies

```yaml
{"type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "id": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}, "policies": {"description": "Access policy for the membership", "type": "array", "items": {"$ref": "#/components/schemas/iam_list_member_policy"}}, "roles": {"description": "Roles assigned to this Member.", "type": "array", "items": {"$ref": "#/components/schemas/iam_role"}}, "status": {"description": "A member's status in the account.", "example": "accepted", "enum": ["accepted", "pending"], "readOnly": true, "x-auditable": true}, "user": {"description": "Details of the user associated to the membership.", "type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "first_name": {"$ref": "#/components/schemas/iam_first_name"}, "id": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}, "last_name": {"$ref": "#/components/schemas/iam_last_name"}, "two_factor_authentication_enabled": {"$ref": "#/components/schemas/iam_two_factor_authentication_enabled"}}, "readOnly": true, "required": ["email"]}}}
```
