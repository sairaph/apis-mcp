---
title: digital-experience-monitoring_device_status_over_time_result
page_id: schema-digital-experience-monitoring-device-status-over-time-result-afb75800
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_device_status_over_time_result

```yaml
{"type": "object", "properties": {"over_time": {"$ref": "#/components/schemas/digital-experience-monitoring_device_state_over_time"}, "top_networks": {"description": "Top networks observed for the device.", "type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_device_state_top_networks_summary"}}}, "required": ["top_networks", "over_time"]}
```
