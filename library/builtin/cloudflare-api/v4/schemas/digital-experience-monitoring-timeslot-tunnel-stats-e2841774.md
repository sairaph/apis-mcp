---
title: digital-experience-monitoring_timeslot_tunnel_stats
page_id: schema-digital-experience-monitoring-timeslot-tunnel-stats-e2841774
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_timeslot_tunnel_stats

```yaml
{"type": "object", "properties": {"timestamp": {"description": "Timestamp of the time slot.", "type": "string", "example": "2023-07-16 15:00:00+00"}, "value": {"description": "WARP tunnel packet and byte counters measured within the stats window for the time slot.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_tunnel_stats"}]}}, "required": ["timestamp", "value"]}
```
