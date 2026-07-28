---
title: RankingsDailyItem
page_id: schema-rankingsdailyitem-34631f3b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RankingsDailyItem

```yaml
{"example": {"date": "2026-05-11", "model_permaslug": "openai/gpt-4o-2024-05-13", "total_tokens": "12345678"}, "properties": {"date": {"description": "UTC calendar date the row is aggregated over (YYYY-MM-DD).", "example": "2026-05-11", "type": "string"}, "model_permaslug": {"description": "Model variant permaslug (e.g. `openai/gpt-4o-2024-05-13`, `openai/gpt-4o-2024-05-13:free`). Non-default variants include a `:variant` suffix and are ranked as their own entry. The reserved value `other` denotes the aggregated row covering every model outside the daily top 50 for that date — always sorted last within its date.", "example": "openai/gpt-4o-2024-05-13", "type": "string"}, "total_tokens": {"description": "Sum of `prompt_tokens + completion_tokens` for the day, returned as a decimal string so 64-bit values are not truncated.", "example": "12345678", "type": "string"}}, "required": ["date", "model_permaslug", "total_tokens"], "type": "object"}
```
