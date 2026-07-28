---
title: digital-experience-monitoring_commands_devices_response
page_id: schema-digital-experience-monitoring-commands-devices-response-2119de0e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_commands_devices_response

```yaml
{"type": "object", "properties": {"devices": {"description": "List of eligible devices", "type": "array", "items": {"properties": {"deviceId": {"description": "Device identifier (UUID v4)", "type": "string"}, "deviceName": {"description": "Device identifier (human readable)", "type": "string"}, "eligible": {"description": "Whether the device is eligible for remote captures", "type": "boolean"}, "ineligibleReason": {"description": "If the device is not eligible, the reason why.", "type": "string"}, "personEmail": {"description": "User contact email address", "type": "string"}, "platform": {"$ref": "#/components/schemas/digital-experience-monitoring_platform"}, "registrationId": {"description": "Device registration identifier (UUID v4). On multi-user devices, this uniquely identifies a user's registration on the device.", "type": "string"}, "status": {"$ref": "#/components/schemas/digital-experience-monitoring_status"}, "timestamp": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}, "version": {"$ref": "#/components/schemas/digital-experience-monitoring_version"}}, "type": "object"}}}}
```
