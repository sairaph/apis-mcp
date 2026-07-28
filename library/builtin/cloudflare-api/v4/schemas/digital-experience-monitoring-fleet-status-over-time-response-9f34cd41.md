---
title: digital-experience-monitoring_fleet_status_over_time_response
page_id: schema-digital-experience-monitoring-fleet-status-over-time-response-9f34cd41
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_fleet_status_over_time_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deviceStats": {"type": "object", "properties": {"byMode": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_schemas-aggregate_stat"}}, "byStatus": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_schemas-aggregate_stat"}}, "uniqueDevicesTotal": {"$ref": "#/components/schemas/digital-experience-monitoring_uniqueDevicesTotal"}}}}}}, "type": "object"}]}
```
