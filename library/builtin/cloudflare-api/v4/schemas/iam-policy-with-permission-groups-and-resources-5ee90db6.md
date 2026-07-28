---
title: iam_policy_with_permission_groups_and_resources
page_id: schema-iam-policy-with-permission-groups-and-resources-5ee90db6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_policy_with_permission_groups_and_resources

```yaml
{"type": "object", "properties": {"effect": {"$ref": "#/components/schemas/iam_effect"}, "id": {"$ref": "#/components/schemas/iam_policy_identifier"}, "permission_groups": {"$ref": "#/components/schemas/iam_permission_groups"}, "resources": {"$ref": "#/components/schemas/iam_resources"}}, "required": ["id", "effect", "permission_groups", "resources"], "title": "Policy with Permission Groups and Resources"}
```
