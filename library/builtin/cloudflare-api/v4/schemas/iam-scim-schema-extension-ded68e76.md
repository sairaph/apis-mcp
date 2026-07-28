---
title: iam_scim_schema_extension
page_id: schema-iam-scim-schema-extension-ded68e76
path: schemas
description: An extension schema associated with a resource type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_schema_extension

An extension schema associated with a resource type.

```yaml
{"description": "An extension schema associated with a resource type.", "type": "object", "properties": {"required": {"description": "Whether the extension is required.", "type": "boolean"}, "schema": {"description": "The URI of the extension schema.", "type": "string"}}, "required": ["schema", "required"], "title": "SCIM Schema Extension"}
```
