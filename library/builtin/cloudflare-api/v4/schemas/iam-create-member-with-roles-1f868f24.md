---
title: iam_create-member-with-roles
page_id: schema-iam-create-member-with-roles-1f868f24
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create-member-with-roles

```yaml
{"type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "roles": {"description": "Array of roles associated with this member.", "type": "array", "items": {"$ref": "#/components/schemas/iam_role_components-schemas-identifier"}}, "status": {"$ref": "#/components/schemas/iam_member-invitation-status"}}, "required": ["email", "roles"], "title": "Add Member with Account Roles"}
```
