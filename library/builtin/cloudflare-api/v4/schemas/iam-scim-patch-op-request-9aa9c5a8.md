---
title: iam_scim_patch_op_request
page_id: schema-iam-scim-patch-op-request-9aa9c5a8
path: schemas
description: Request body for a SCIM PATCH operation on a User resource (RFC 7644 Section 3.5.2).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_patch_op_request

Request body for a SCIM PATCH operation on a User resource (RFC 7644 Section 3.5.2).

```yaml
{"description": "Request body for a SCIM PATCH operation on a User resource (RFC 7644 Section 3.5.2).\n", "type": "object", "properties": {"Operations": {"description": "List of PATCH operations to apply.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_patch_operation"}, "example": [{"op": "replace", "value": {"active": false}}]}, "schemas": {"description": "Must contain `urn:ietf:params:scim:api:messages:2.0:PatchOp`.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "required": ["schemas", "Operations"], "title": "SCIM PatchOp Request (Users)"}
```
