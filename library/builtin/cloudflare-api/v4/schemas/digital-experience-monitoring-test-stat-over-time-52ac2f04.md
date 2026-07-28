---
title: digital-experience-monitoring_test_stat_over_time
page_id: schema-digital-experience-monitoring-test-stat-over-time-52ac2f04
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_test_stat_over_time

```yaml
{"type": "object", "properties": {"avg": {"description": "average observed in the time period.", "type": "integer", "nullable": true}, "max": {"description": "highest observed in the time period.", "type": "integer", "nullable": true}, "min": {"description": "lowest observed in the time period.", "type": "integer", "nullable": true}, "slots": {"type": "array", "items": {"properties": {"timestamp": {"type": "string", "example": "2023-07-16 15:00:00+00"}, "value": {"type": "integer"}}, "required": ["timestamp", "value"], "type": "object"}}}, "required": ["slots"]}
```
