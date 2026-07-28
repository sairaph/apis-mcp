---
title: teams-devices_policy_summary
page_id: schema-teams-devices-policy-summary-c9a28a1a
path: schemas
description: The device settings profile assigned to this registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_policy_summary

The device settings profile assigned to this registration.

```yaml
{"description": "The device settings profile assigned to this registration.", "type": "object", "properties": {"default": {"description": "Whether the device settings profile is the default profile for the account.", "type": "boolean"}, "deleted": {"description": "Whether the device settings profile was deleted.", "type": "boolean"}, "id": {"description": "The ID of the device settings profile.", "type": "string", "example": "11ffb86f-3f0c-4306-b4a2-e62f872b166a"}, "name": {"description": "The name of the device settings profile.", "type": "string"}, "updated_at": {"description": "The RFC3339 timestamp of when the device settings profile last changed for the registration.", "type": "string", "example": "2025-02-14T13:17:00Z"}}, "required": ["id", "name", "updated_at", "default", "deleted"]}
```
