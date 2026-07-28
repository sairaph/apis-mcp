---
title: digital-experience-monitoring_warp_config_change_event
page_id: schema-digital-experience-monitoring-warp-config-change-event-cd5ec090
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_warp_config_change_event

```yaml
{"type": "object", "properties": {"device_id": {"description": "The device ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}]}, "device_registration": {"description": "Deprecated: use registration_id. The device registration ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}], "deprecated": true, "x-stainless-deprecation-message": "Use `registration_id` instead."}, "from": {"description": "The details for the WARP configuration that was switched from.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_warp_config_details"}]}, "hostname": {"description": "The hostname of the machine the event is from.", "type": "string"}, "registration_id": {"description": "The device registration ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}]}, "serial_number": {"description": "The serial number of the machine the event is from.", "type": "string"}, "timestamp": {"description": "The event time.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_timestamp_datetime"}]}, "to": {"description": "The details for the WARP configuration that was switched to.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_warp_config_details"}]}, "user_email": {"description": "Email tied to the device.", "type": "string"}}}
```
