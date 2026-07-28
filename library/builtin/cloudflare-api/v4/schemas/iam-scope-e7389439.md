---
title: iam_scope
page_id: schema-iam-scope-e7389439
path: schemas
description: A scope is a combination of scope objects which provides additional context.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scope

A scope is a combination of scope objects which provides additional context.

```yaml
{"description": "A scope is a combination of scope objects which provides additional context.", "type": "object", "properties": {"key": {"$ref": "#/components/schemas/iam_scope_key"}, "objects": {"description": "A list of scope objects for additional context.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scope_object"}}}, "required": ["key", "objects"]}
```
