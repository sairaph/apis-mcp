---
title: art_TimeseriesResult
page_id: schema-art-timeseriesresult-d6aa288d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_TimeseriesResult

```yaml
{"type": "object", "properties": {"resolution": {"description": "The resolution used for time bucketing.", "type": "string", "example": "hour"}, "slots": {"description": "Time-bucketed result rows. Each slot contains a `time_bucket` field plus the requested stats and group-by dimensions.\n", "type": "array", "items": {"$ref": "#/components/schemas/art_ResultValues"}, "example": [{"attemptsTotal": 1234, "country": "US", "time_bucket": "2024-11-05T00:00:00Z"}]}}, "required": ["slots", "resolution"]}
```
