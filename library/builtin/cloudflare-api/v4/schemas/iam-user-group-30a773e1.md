---
title: iam_user_group
page_id: schema-iam-user-group-30a773e1
path: schemas
description: A group of policies resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_user_group

A group of policies resources.

```yaml
{"description": "A group of policies resources.", "type": "object", "properties": {"created_on": {"description": "Timestamp for the creation of the user group", "type": "string", "format": "date-time", "example": "2024-03-01T12:21:02.0000Z", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/iam_user_group_identifier"}, "modified_on": {"description": "Last time the user group was modified.", "type": "string", "format": "date-time", "example": "2024-03-01T12:21:02.0000Z", "readOnly": true, "x-auditable": true}, "name": {"description": "Name of the user group.", "type": "string", "example": "My New User Group", "readOnly": true, "x-auditable": true}, "policies": {"description": "Policies attached to the User group", "type": "array", "items": {"description": "Policy", "properties": {"access": {"$ref": "#/components/schemas/iam_access"}, "id": {"$ref": "#/components/schemas/iam_policy_identifier"}, "permission_groups": {"$ref": "#/components/schemas/iam_permission_groups"}, "resource_groups": {"$ref": "#/components/schemas/iam_resource_groups"}}, "title": "Authorization Policy"}, "title": "User Group Policies"}}, "required": ["id", "name", "created_on", "modified_on"]}
```
