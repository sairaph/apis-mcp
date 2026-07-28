---
title: access_device_posture_rule
page_id: schema-access-device-posture-rule-5437892a
path: schemas
description: Enforces a device posture rule has run successfully
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_device_posture_rule

Enforces a device posture rule has run successfully

```yaml
{"description": "Enforces a device posture rule has run successfully", "type": "object", "properties": {"device_posture": {"type": "object", "properties": {"integration_uid": {"description": "The ID of a device posture integration.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}}, "required": ["integration_uid"]}}, "required": ["device_posture"], "title": "Device Posture"}
```
