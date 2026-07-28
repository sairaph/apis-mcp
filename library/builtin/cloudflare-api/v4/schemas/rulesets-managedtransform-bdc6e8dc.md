---
title: rulesets_ManagedTransform
page_id: schema-rulesets-managedtransform-bdc6e8dc
path: schemas
description: A Managed Transform object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ManagedTransform

A Managed Transform object.

```yaml
{"description": "A Managed Transform object.", "type": "object", "properties": {"conflicts_with": {"description": "The Managed Transforms that this Managed Transform conflicts with.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_ManagedTransformId"}, {"example": "add_true_client_ip_headers"}]}, "minItems": 1, "readOnly": true, "title": "Conflicts With", "uniqueItems": true, "x-stainless-skip": ["terraform"]}, "enabled": {"description": "Whether the Managed Transform is enabled.", "type": "boolean", "example": true, "title": "Enabled"}, "has_conflict": {"description": "Whether the Managed Transform conflicts with the currently-enabled Managed Transforms.", "type": "boolean", "example": false, "readOnly": true, "title": "Has Conflict", "x-stainless-skip": ["terraform"]}, "id": {"$ref": "#/components/schemas/rulesets_ManagedTransformId"}}, "required": ["id", "enabled", "has_conflict"], "title": "Managed Transform"}
```
