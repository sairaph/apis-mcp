---
title: PercentileStats
page_id: schema-percentilestats-fd21e1ba
path: schemas
description: Latency percentiles in milliseconds over the last 30 minutes. Latency measures time to first token. Only visible when authenticated with an API key or cookie; returns null for unauthenticated requests.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PercentileStats

Latency percentiles in milliseconds over the last 30 minutes. Latency measures time to first token. Only visible when authenticated with an API key or cookie; returns null for unauthenticated requests.

```yaml
{"description": "Latency percentiles in milliseconds over the last 30 minutes. Latency measures time to first token. Only visible when authenticated with an API key or cookie; returns null for unauthenticated requests.", "example": {"p50": 25.5, "p75": 35.2, "p90": 48.7, "p99": 85.3}, "properties": {"p50": {"description": "Median (50th percentile)", "example": 25.5, "format": "double", "type": "number"}, "p75": {"description": "75th percentile", "example": 35.2, "format": "double", "type": "number"}, "p90": {"description": "90th percentile", "example": 48.7, "format": "double", "type": "number"}, "p99": {"description": "99th percentile", "example": 85.3, "format": "double", "type": "number"}}, "required": ["p50", "p75", "p90", "p99"], "type": ["object", "null"]}
```
