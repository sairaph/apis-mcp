---
title: one_FeatureScope
page_id: schema-one-featurescope-be270674
path: schemas
description: A feature with its additional scopes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_FeatureScope

A feature with its additional scopes.

```yaml
{"description": "A feature with its additional scopes.", "type": "object", "properties": {"description": {"description": "Feature description.", "type": "string"}, "display_name": {"description": "Human-readable feature name.", "type": "string"}, "id": {"description": "Feature identifier.", "type": "string"}, "scopes": {"description": "Additional scopes when feature is enabled.", "type": "array", "items": {"$ref": "#/components/schemas/one_Permission"}}}, "required": ["description", "display_name", "id", "scopes"]}
```
