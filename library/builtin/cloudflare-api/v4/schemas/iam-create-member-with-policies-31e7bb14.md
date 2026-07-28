---
title: iam_create-member-with-policies
page_id: schema-iam-create-member-with-policies-31e7bb14
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create-member-with-policies

```yaml
{"type": "object", "properties": {"email": {"$ref": "#/components/schemas/iam_email"}, "policies": {"description": "Array of policies associated with this member.", "type": "array", "items": {"$ref": "#/components/schemas/iam_create_member_policy"}}, "status": {"$ref": "#/components/schemas/iam_member-invitation-status"}}, "required": ["email", "policies"], "title": "Add Member with Policies"}
```
