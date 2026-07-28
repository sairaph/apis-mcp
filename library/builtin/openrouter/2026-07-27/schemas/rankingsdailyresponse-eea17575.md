---
title: RankingsDailyResponse
page_id: schema-rankingsdailyresponse-eea17575
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RankingsDailyResponse

```yaml
{"example": {"data": [{"date": "2026-05-11", "model_permaslug": "openai/gpt-4o-2024-05-13", "total_tokens": "12345678"}, {"date": "2026-05-11", "model_permaslug": "anthropic/claude-3.5-sonnet-20241022", "total_tokens": "9876543"}], "meta": {"as_of": "2026-05-12T02:00:00Z", "end_date": "2026-05-11", "start_date": "2026-04-12", "version": "v1"}}, "properties": {"data": {"description": "Up to 51 rows per day — the top 50 public models by `total_tokens` for each UTC calendar date in the window, plus one aggregated `other` row summing every model outside that top 50 (omitted when the long tail is empty). Rows are sorted by `date` ascending, then by `total_tokens` descending, with `other` pinned last within its date. Ties between real models break alphabetically on `model_permaslug` so the order is stable across requests.", "items": {"$ref": "#/components/schemas/RankingsDailyItem"}, "type": "array"}, "meta": {"$ref": "#/components/schemas/RankingsDailyMeta"}}, "required": ["data", "meta"], "type": "object"}
```
