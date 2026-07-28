---
title: iam_scim_patch_operation
page_id: schema-iam-scim-patch-operation-fb4de4f1
path: schemas
description: A single PATCH operation (RFC 7644 Section 3.5.2).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_patch_operation

A single PATCH operation (RFC 7644 Section 3.5.2).

```yaml
{"description": "A single PATCH operation (RFC 7644 Section 3.5.2).", "type": "object", "properties": {"op": {"description": "The operation to perform. Only `replace` is currently supported; `add` and `remove` are accepted without error but have no effect. Matched case-insensitively.\n", "type": "string", "example": "replace", "enum": ["add", "remove", "replace"], "x-auditable": true}, "path": {"description": "Attribute path targeted by this operation. When absent, `value` must be a singular complex attribute.\n", "type": "string", "example": "active", "x-auditable": true}, "value": {"description": "The value(s) for the operation. For `replace` without a path, this should be an object of attribute name/value pairs. For member path operations, this should be an array of member value objects.\n", "example": {"active": false}}}, "required": ["op"], "title": "SCIM Patch Operation"}
```
