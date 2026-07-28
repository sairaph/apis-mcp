---
title: PercentileThroughputCutoffs
page_id: schema-percentilethroughputcutoffs-528c3147
path: schemas
description: Percentile-based throughput cutoffs. All specified cutoffs must be met for an endpoint to be preferred.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PercentileThroughputCutoffs

Percentile-based throughput cutoffs. All specified cutoffs must be met for an endpoint to be preferred.

```yaml
{"description": "Percentile-based throughput cutoffs. All specified cutoffs must be met for an endpoint to be preferred.", "example": {"p50": 100, "p90": 50}, "properties": {"p50": {"description": "Minimum p50 throughput (tokens/sec)", "format": "double", "type": ["number", "null"]}, "p75": {"description": "Minimum p75 throughput (tokens/sec)", "format": "double", "type": ["number", "null"]}, "p90": {"description": "Minimum p90 throughput (tokens/sec)", "format": "double", "type": ["number", "null"]}, "p99": {"description": "Minimum p99 throughput (tokens/sec)", "format": "double", "type": ["number", "null"]}}, "type": "object"}
```
