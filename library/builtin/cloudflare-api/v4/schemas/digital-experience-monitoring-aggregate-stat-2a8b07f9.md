---
title: digital-experience-monitoring_aggregate_stat
page_id: schema-digital-experience-monitoring-aggregate-stat-2a8b07f9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_aggregate_stat

```yaml
{"type": "object", "properties": {"avgMs": {"type": "integer", "nullable": true}, "deltaPct": {"type": "number", "format": "float", "nullable": true}, "timePeriod": {"$ref": "#/components/schemas/digital-experience-monitoring_aggregate_time_period"}}, "required": ["timePeriod"]}
```
