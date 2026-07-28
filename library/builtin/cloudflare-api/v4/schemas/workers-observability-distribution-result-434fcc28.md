---
title: workers-observability_distribution_result
page_id: schema-workers-observability-distribution-result-434fcc28
path: schemas
description: Bucketed 2D histogram of a numeric field over time. Present when chartType is 'distribution'.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-observability_distribution_result

Bucketed 2D histogram of a numeric field over time. Present when chartType is 'distribution'.

```yaml
{"description": "Bucketed 2D histogram of a numeric field over time. Present when chartType is 'distribution'.", "type": "object", "properties": {"bins": {"description": "Time-bucket labels (ISO-8601 strings), one per matrix column.", "type": "array", "items": {"type": "string"}}, "bucketBoundaries": {"description": "Raw bucket edges in the value's native unit, length buckets.length + 1. Used for the colour scale and percentile mapping.", "type": "array", "items": {"type": "number"}}, "bucketMode": {"description": "Bucketing scheme used to derive the boundaries. 'log' produces geometric edges; 'linear' produces fixed-width edges.", "type": "string", "enum": ["log", "linear"]}, "buckets": {"description": "Value-range labels, one per matrix row (e.g. '50–100ms').", "type": "array", "items": {"type": "string"}}, "matrix": {"description": "Sampling-corrected counts. matrix[bucketIdx][binIdx] is the estimated number of events in value-bucket 'bucketIdx' during time-bin 'binIdx'.", "type": "array", "items": {"items": {"type": "number"}, "type": "array"}}}, "required": ["bins", "buckets", "matrix", "bucketBoundaries", "bucketMode"]}
```
