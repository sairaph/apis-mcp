---
title: digital-experience-monitoring_fleet_status_live_response
page_id: schema-digital-experience-monitoring-fleet-status-live-response-6cd36dc0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_fleet_status_live_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deviceStats": {"type": "object", "properties": {"byColo": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_live_stat"}, "nullable": true}, "byMode": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_live_stat"}, "nullable": true}, "byPlatform": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_live_stat"}, "nullable": true}, "byStatus": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_live_stat"}, "nullable": true}, "byVersion": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_live_stat"}, "nullable": true}, "uniqueDevicesTotal": {"$ref": "#/components/schemas/digital-experience-monitoring_uniqueDevicesTotal"}}}}}}, "type": "object"}]}
```
