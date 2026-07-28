---
title: digital-experience-monitoring_warp_toggle_change_event
page_id: schema-digital-experience-monitoring-warp-toggle-change-event-7d43d05d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_warp_toggle_change_event

```yaml
{"type": "object", "properties": {"account_name": {"description": "The account name.", "type": "string"}, "account_tag": {"description": "The public account identifier.", "type": "string"}, "device_id": {"description": "The device ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}]}, "device_registration": {"description": "Deprecated: use registration_id. The device registration ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}], "deprecated": true, "x-stainless-deprecation-message": "Use `registration_id` instead."}, "hostname": {"description": "The hostname of the machine the event is from.", "type": "string"}, "registration_id": {"description": "The device registration ID.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}]}, "serial_number": {"description": "The serial number of the machine the event is from.", "type": "string"}, "timestamp": {"description": "The event time.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_timestamp_datetime"}]}, "toggle": {"description": "The state of the WARP toggle.", "type": "string", "enum": ["on", "off"]}, "user_email": {"description": "Email tied to the device.", "type": "string"}}}
```
