---
title: TaskClassificationResponse
page_id: schema-taskclassificationresponse-b3afd36b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TaskClassificationResponse

```yaml
{"example": {"data": {"as_of": "2026-06-17", "classifications": [{"category_token_share": 0.48, "category_usage_share": 0.51, "display_name": "Code Generation", "macro_category": "code", "models": [{"id": "openai/gpt-4.1-mini", "tag_token_share": 0.75, "tag_usage_share": 0.55}], "tag": "code:general_impl", "token_share": 0.31, "usage_share": 0.23}], "macro_categories": [{"key": "code", "label": "Code", "token_share": 0.52, "usage_share": 0.45}], "window_days": 7}}, "properties": {"data": {"properties": {"as_of": {"description": "UTC date (YYYY-MM-DD) of the window upper bound (yesterday). Data is exclusive of the current incomplete UTC day. This is the expected latest date in the snapshot; it does not confirm data presence for that date.", "example": "2026-06-17", "type": "string"}, "classifications": {"description": "Per-task classification market-share data, sorted by usage_share descending.", "items": {"$ref": "#/components/schemas/TaskClassificationItem"}, "type": "array"}, "macro_categories": {"description": "Aggregate market-share data per macro-category (code, data, agent, general).", "items": {"$ref": "#/components/schemas/TaskClassificationMacroCategory"}, "type": "array"}, "window_days": {"description": "Number of trailing days covered by this snapshot.", "example": 7, "type": "integer"}}, "required": ["window_days", "as_of", "classifications", "macro_categories"], "type": "object"}}, "required": ["data"], "type": "object"}
```
