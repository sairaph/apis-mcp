---
title: teams-devices_device-posture-rules
page_id: schema-teams-devices-device-posture-rules-9f89b1d8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_device-posture-rules

```yaml
{"type": "object", "properties": {"description": {"$ref": "#/components/schemas/teams-devices_description"}, "enabled": {"description": "Whether the rule is enabled. This is a computed, read-only value. It is false for deprecated Kolide posture rules that still use the issue_count input, and true otherwise.", "type": "boolean", "example": true, "readOnly": true}, "expiration": {"$ref": "#/components/schemas/teams-devices_expiration"}, "id": {"$ref": "#/components/schemas/teams-devices_uuid"}, "input": {"$ref": "#/components/schemas/teams-devices_input"}, "match": {"$ref": "#/components/schemas/teams-devices_match"}, "name": {"$ref": "#/components/schemas/teams-devices_name"}, "schedule": {"$ref": "#/components/schemas/teams-devices_schedule"}, "type": {"$ref": "#/components/schemas/teams-devices_type"}}}
```
