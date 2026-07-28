---
title: PercentileLatencyCutoffs
page_id: schema-percentilelatencycutoffs-a5bd74d4
path: schemas
description: Percentile-based latency cutoffs. All specified cutoffs must be met for an endpoint to be preferred.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PercentileLatencyCutoffs

Percentile-based latency cutoffs. All specified cutoffs must be met for an endpoint to be preferred.

```yaml
{"description": "Percentile-based latency cutoffs. All specified cutoffs must be met for an endpoint to be preferred.", "example": {"p50": 5, "p90": 10}, "properties": {"p50": {"description": "Maximum p50 latency (seconds)", "format": "double", "type": ["number", "null"]}, "p75": {"description": "Maximum p75 latency (seconds)", "format": "double", "type": ["number", "null"]}, "p90": {"description": "Maximum p90 latency (seconds)", "format": "double", "type": ["number", "null"]}, "p99": {"description": "Maximum p99 latency (seconds)", "format": "double", "type": ["number", "null"]}}, "type": "object"}
```
