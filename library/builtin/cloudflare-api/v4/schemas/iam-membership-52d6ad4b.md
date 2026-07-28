---
title: iam_membership
page_id: schema-iam-membership-52d6ad4b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_membership

```yaml
{"type": "object", "properties": {"account": {"$ref": "#/components/schemas/iam_schemas-account"}, "api_access_enabled": {"$ref": "#/components/schemas/iam_api_access_enabled"}, "id": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}, "permissions": {"description": "All access permissions for the user at the account.", "allOf": [{"$ref": "#/components/schemas/iam_permissions"}], "readOnly": true}, "roles": {"$ref": "#/components/schemas/iam_role_names"}, "status": {"$ref": "#/components/schemas/iam_schemas-status"}}}
```
