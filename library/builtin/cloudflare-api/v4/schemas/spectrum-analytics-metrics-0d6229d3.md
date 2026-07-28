---
title: spectrum-analytics_metrics
page_id: schema-spectrum-analytics-metrics-0d6229d3
path: schemas
description: |-
    One or more metrics to compute. Options are:

    Metric                    | Name                                | Example                  | Unit
    --------------------------|-------------------------------------|--------------------------|--------------------------
    count                     | Count of total events               | 1000                     | Count
    bytesIngress              | Sum of ingress bytes                | 1000                     | Sum
    bytesEgress               | Sum of egress bytes                 | 1000                     | Sum
    durationAvg               | Average connection duration         | 1.0                      | Time in milliseconds
    durationMedian            | Median connection duration          | 1.0                      | Time in milliseconds
    duration90th              | 90th percentile connection duration | 1.0                      | Time in milliseconds
    duration99th              | 99th percentile connection duration | 1.0                      | Time in milliseconds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_metrics

One or more metrics to compute. Options are:

Metric                    | Name                                | Example                  | Unit
--------------------------|-------------------------------------|--------------------------|--------------------------
count                     | Count of total events               | 1000                     | Count
bytesIngress              | Sum of ingress bytes                | 1000                     | Sum
bytesEgress               | Sum of egress bytes                 | 1000                     | Sum
durationAvg               | Average connection duration         | 1.0                      | Time in milliseconds
durationMedian            | Median connection duration          | 1.0                      | Time in milliseconds
duration90th              | 90th percentile connection duration | 1.0                      | Time in milliseconds
duration99th              | 99th percentile connection duration | 1.0                      | Time in milliseconds.

```yaml
{"description": "One or more metrics to compute. Options are:\n\nMetric                    | Name                                | Example                  | Unit\n--------------------------|-------------------------------------|--------------------------|--------------------------\ncount                     | Count of total events               | 1000                     | Count\nbytesIngress              | Sum of ingress bytes                | 1000                     | Sum\nbytesEgress               | Sum of egress bytes                 | 1000                     | Sum\ndurationAvg               | Average connection duration         | 1.0                      | Time in milliseconds\ndurationMedian            | Median connection duration          | 1.0                      | Time in milliseconds\nduration90th              | 90th percentile connection duration | 1.0                      | Time in milliseconds\nduration99th              | 99th percentile connection duration | 1.0                      | Time in milliseconds.", "type": "array", "items": {"enum": ["count", "bytesIngress", "bytesEgress", "durationAvg", "durationMedian", "duration90th", "duration99th"], "type": "string"}, "example": ["count", "bytesIngress"]}
```
