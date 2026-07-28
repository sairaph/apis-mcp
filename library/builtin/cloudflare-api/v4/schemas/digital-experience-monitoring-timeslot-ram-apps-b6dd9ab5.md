---
title: digital-experience-monitoring_timeslot_ram_apps
page_id: schema-digital-experience-monitoring-timeslot-ram-apps-b6dd9ab5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_timeslot_ram_apps

```yaml
{"type": "object", "properties": {"timestamp": {"description": "Timestamp of the time slot.", "type": "string", "example": "2023-07-16 15:00:00+00"}, "value": {"description": "Top RAM-consuming applications for the time slot.", "type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_ram_used_pct_by_app"}}}, "required": ["timestamp", "value"]}
```
