---
title: iam_scim_group_patch_operation
page_id: schema-iam-scim-group-patch-operation-17fd715b
path: schemas
description: A single PATCH operation for a Group resource. Supports `add`, `remove`, and `replace` on `members`, `displayName`, and `externalId`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_group_patch_operation

A single PATCH operation for a Group resource. Supports `add`, `remove`, and `replace` on `members`, `displayName`, and `externalId`.

```yaml
{"description": "A single PATCH operation for a Group resource. Supports `add`, `remove`, and `replace` on `members`, `displayName`, and `externalId`.\n", "type": "object", "properties": {"op": {"description": "The operation to perform.", "type": "string", "example": "add", "enum": ["add", "remove", "replace"], "x-auditable": true}, "path": {"description": "Attribute path targeted by this operation. Use `members` to modify group membership. May also include a filter expression to target specific members, e.g. `members[value eq \"userTag\"]`.\n", "type": "string", "example": "members", "x-auditable": true}, "value": {"description": "The value(s) for the operation. For member add/replace operations, an array of member value objects. For `displayName` or `externalId` updates, a string value.\n", "example": [{"value": "023e105f4ecef8ad9ca31a8372d0c353"}], "oneOf": [{"items": {"$ref": "#/components/schemas/iam_scim_group_patch_member_value"}, "type": "array"}, {"type": "string"}]}}, "required": ["op"], "title": "SCIM Group Patch Operation"}
```
