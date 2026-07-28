---
title: iam_resource_group
page_id: schema-iam-resource-group-79ff81f2
path: schemas
description: A group of scoped resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_resource_group

A group of scoped resources.

```yaml
{"description": "A group of scoped resources.", "type": "object", "properties": {"id": {"description": "Identifier of the resource group.", "type": "string", "example": "6d7f2f5f5b1d4a0e9081fdc98d432fd1", "readOnly": true, "x-auditable": true}, "meta": {"description": "Attributes associated to the resource group.", "type": "object", "example": {"editable": "false"}, "properties": {"key": {"type": "string"}, "value": {"type": "string"}}}, "name": {"description": "Name of the resource group.", "type": "string", "example": "com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4", "readOnly": true, "x-auditable": true}, "scope": {"$ref": "#/components/schemas/iam_scope"}}, "required": ["id", "scope"]}
```
