---
title: iam_organization
page_id: schema-iam-organization-10c7773d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_organization

```yaml
{"type": "object", "properties": {"id": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}, "name": {"$ref": "#/components/schemas/iam_schemas-name"}, "permissions": {"$ref": "#/components/schemas/iam_schemas-permissions"}, "roles": {"description": "List of roles that a user has within an organization.", "type": "array", "items": {"example": "All Privileges - Super Administrator", "maxLength": 120, "type": "string", "x-auditable": true}, "readOnly": true}, "status": {"$ref": "#/components/schemas/iam_components-schemas-status"}}}
```
