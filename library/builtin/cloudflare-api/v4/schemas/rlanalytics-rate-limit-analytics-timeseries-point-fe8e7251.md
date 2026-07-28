---
title: rlanalytics_rate_limit_analytics_timeseries_point
page_id: schema-rlanalytics-rate-limit-analytics-timeseries-point-fe8e7251
path: schemas
description: Analytics data for a single time segment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rlanalytics_rate_limit_analytics_timeseries_point

Analytics data for a single time segment.

```yaml
{"description": "Analytics data for a single time segment.", "type": "object", "properties": {"rules": {"description": "Contains rule-level analytics for this time segment.", "type": "object", "additionalProperties": {"$ref": "#/components/schemas/rlanalytics_rate_limit_analytics_rule_rollup"}}, "since": {"description": "Start of the time segment formatted as RFC 3339.", "type": "string", "format": "date-time"}, "until": {"description": "Exclusive end of the time segment formatted as RFC 3339.", "type": "string", "format": "date-time"}}, "required": ["since", "until", "rules"]}
```
