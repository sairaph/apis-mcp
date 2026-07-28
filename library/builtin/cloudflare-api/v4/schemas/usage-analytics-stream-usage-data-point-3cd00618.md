---
title: usage-analytics_stream_usage_data_point
page_id: schema-usage-analytics-stream-usage-data-point-3cd00618
path: schemas
description: A single Stream usage data point for a time period.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# usage-analytics_stream_usage_data_point

A single Stream usage data point for a time period.

```yaml
{"description": "A single Stream usage data point for a time period.", "type": "object", "properties": {"streamMinutesViewed": {"description": "Number of Stream billable minutes viewed in this time period.", "type": "integer", "format": "int64", "example": 12500}, "ts": {"description": "Unix timestamp (epoch seconds) for the start of this time period.", "type": "integer", "format": "int64", "example": 1693526400}}}
```
