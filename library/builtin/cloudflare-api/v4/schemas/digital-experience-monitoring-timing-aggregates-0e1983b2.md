---
title: digital-experience-monitoring_timing_aggregates
page_id: schema-digital-experience-monitoring-timing-aggregates-0e1983b2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_timing_aggregates

```yaml
{"type": "object", "properties": {"avgMs": {"type": "integer", "nullable": true}, "history": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_aggregate_stat"}}, "overTime": {"type": "object", "nullable": true, "properties": {"timePeriod": {"$ref": "#/components/schemas/digital-experience-monitoring_aggregate_time_period"}, "values": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_aggregate_time_slot"}}}, "required": ["values", "timePeriod"]}}, "required": ["history"]}
```
