---
title: TaskClassificationItem
page_id: schema-taskclassificationitem-41ce8d44
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TaskClassificationItem

```yaml
{"example": {"category_token_share": 0.48, "category_usage_share": 0.51, "display_name": "Code Generation", "macro_category": "code", "models": [{"id": "openai/gpt-4.1-mini", "tag_token_share": 0.75, "tag_usage_share": 0.55}, {"id": "anthropic/claude-sonnet-4", "tag_token_share": 0.12, "tag_usage_share": 0.2}], "tag": "code:general_impl", "token_share": 0.31, "usage_share": 0.23}, "properties": {"category_token_share": {"description": "Fraction of this classification's token volume within its macro-category (0–1). Sums to 1 across all classifications sharing the same `macro_category`.", "example": 0.48, "format": "double", "type": "number"}, "category_usage_share": {"description": "Fraction of this classification's usage within its macro-category (0–1). Sums to 1 across all classifications sharing the same `macro_category`.", "example": 0.51, "format": "double", "type": "number"}, "display_name": {"description": "Human-readable label for the classification.", "example": "Code Generation", "type": "string"}, "macro_category": {"description": "Coarse grouping derived from the tag prefix: `code`, `data`, `agent`, or `general`.", "example": "code", "type": "string"}, "models": {"description": "Top models for this classification by request volume, sorted descending. Each entry reports the model's share of this classification's requests and tokens.", "items": {"$ref": "#/components/schemas/TaskClassificationModel"}, "type": "array"}, "tag": {"description": "Classification tag identifier (e.g. `code:general_impl`, `agent:web_search`).", "example": "code:general_impl", "type": "string"}, "token_share": {"description": "Fraction of classified sampled token volume (prompt + completion) attributed to this classification (0–1). The unclassified `other` bucket is excluded from the denominator.", "example": 0.31, "format": "double", "type": "number"}, "usage_share": {"description": "Fraction of classified sampled requests attributed to this classification (0–1). The unclassified `other` bucket is excluded from the denominator.", "example": 0.23, "format": "double", "type": "number"}}, "required": ["tag", "display_name", "macro_category", "usage_share", "token_share", "category_usage_share", "category_token_share", "models"], "type": "object"}
```
