---
title: iam_create-scope
page_id: schema-iam-create-scope-d63f0a19
path: schemas
description: A scope is a combination of scope objects which provides additional context.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create-scope

A scope is a combination of scope objects which provides additional context.

```yaml
{"description": "A scope is a combination of scope objects which provides additional context.", "type": "object", "properties": {"key": {"$ref": "#/components/schemas/iam_create_resource_group_scope_scope_key"}, "objects": {"description": "A list of scope objects for additional context. The number of Scope objects should not be zero.", "type": "array", "items": {"$ref": "#/components/schemas/iam_create_resource_group_scope_scope_object"}}}, "required": ["key", "objects"]}
```
