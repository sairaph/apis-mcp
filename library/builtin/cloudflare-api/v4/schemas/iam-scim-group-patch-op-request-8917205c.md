---
title: iam_scim_group_patch_op_request
page_id: schema-iam-scim-group-patch-op-request-8917205c
path: schemas
description: Request body for a SCIM PATCH operation on a Group resource (RFC 7644 Section 3.5.2).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_group_patch_op_request

Request body for a SCIM PATCH operation on a Group resource (RFC 7644 Section 3.5.2).

```yaml
{"description": "Request body for a SCIM PATCH operation on a Group resource (RFC 7644 Section 3.5.2).\n", "type": "object", "properties": {"Operations": {"description": "List of PATCH operations to apply.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_group_patch_operation"}, "example": [{"op": "add", "path": "members", "value": [{"value": "023e105f4ecef8ad9ca31a8372d0c353"}]}]}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "required": ["schemas", "Operations"], "title": "SCIM PatchOp Request (Groups)"}
```
