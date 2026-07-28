---
title: iam_role
page_id: schema-iam-role-db6f1e55
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_role

```yaml
{"type": "object", "properties": {"description": {"description": "Description of role's permissions.", "type": "string", "example": "Administrative access to the entire Account", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/iam_role_components-schemas-identifier"}, "name": {"description": "Role name.", "type": "string", "example": "Account Administrator", "maxLength": 120, "readOnly": true, "x-auditable": true}, "permissions": {"allOf": [{"$ref": "#/components/schemas/iam_permissions"}, {"readOnly": true}]}}, "required": ["id", "name", "description", "permissions"]}
```
