---
title: iam_update-member-with-policies
page_id: schema-iam-update-member-with-policies-91975ddd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_update-member-with-policies

```yaml
{"type": "object", "properties": {"policies": {"description": "Array of policies associated with this member.", "type": "array", "items": {"$ref": "#/components/schemas/iam_create_member_policy"}}}, "required": ["policies"], "title": "Update Member with Policies"}
```
