---
title: RankingsDailyMeta
page_id: schema-rankingsdailymeta-b2a06060
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RankingsDailyMeta

```yaml
{"example": {"as_of": "2026-05-12T02:00:00Z", "end_date": "2026-05-11", "start_date": "2026-04-12", "version": "v1"}, "properties": {"as_of": {"description": "ISO-8601 timestamp of when the response was generated. Reflects data-freshness because the underlying materialized view continuously ingests upstream events.", "example": "2026-05-12T02:00:00Z", "type": "string"}, "end_date": {"description": "Resolved end of the date window (UTC, inclusive).", "example": "2026-05-11", "type": "string"}, "start_date": {"description": "Resolved start of the date window (UTC, inclusive).", "example": "2026-04-12", "type": "string"}, "version": {"description": "Dataset version. Field names and grain are stable for the life of `v1`.", "enum": ["v1"], "type": "string"}}, "required": ["as_of", "version", "start_date", "end_date"], "type": "object"}
```
