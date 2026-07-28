---
title: AppRankingsResponse
page_id: schema-apprankingsresponse-e90b9b5b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AppRankingsResponse

```yaml
{"example": {"data": [{"app_id": 12345, "app_name": "Cline", "rank": 1, "total_requests": 4321, "total_tokens": "12345678"}, {"app_id": 67890, "app_name": "Roo Code", "rank": 2, "total_requests": 2109, "total_tokens": "9876543"}], "meta": {"as_of": "2026-05-12T02:00:00Z", "end_date": "2026-05-11", "start_date": "2026-04-12", "version": "v1"}}, "properties": {"data": {"description": "Apps ranked per the requested `sort`, re-numbered 1..N after category filtering. `popular` sorts by `total_tokens` descending; `trending` sorts by absolute excess token growth descending and may return fewer than `limit` rows when few apps are growing.", "items": {"$ref": "#/components/schemas/AppRankingsItem"}, "type": "array"}, "meta": {"$ref": "#/components/schemas/RankingsDailyMeta"}}, "required": ["data", "meta"], "type": "object"}
```
