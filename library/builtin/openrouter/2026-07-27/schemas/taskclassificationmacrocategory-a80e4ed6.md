---
title: TaskClassificationMacroCategory
page_id: schema-taskclassificationmacrocategory-a80e4ed6
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TaskClassificationMacroCategory

```yaml
{"example": {"key": "code", "label": "Code", "token_share": 0.52, "usage_share": 0.45}, "properties": {"key": {"description": "Macro-category identifier.", "example": "code", "type": "string"}, "label": {"description": "Human-readable label for the macro-category.", "example": "Code", "type": "string"}, "token_share": {"description": "Combined token share of all classifications in this macro-category (0–1).", "example": 0.52, "format": "double", "type": "number"}, "usage_share": {"description": "Combined usage share of all classifications in this macro-category (0–1).", "example": 0.45, "format": "double", "type": "number"}}, "required": ["key", "label", "usage_share", "token_share"], "type": "object"}
```
