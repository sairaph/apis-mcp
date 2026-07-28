---
title: teams-devices_registration_device_details
page_id: schema-teams-devices-registration-device-details-60112181
path: schemas
description: Device details embedded inside of a registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_registration_device_details

Device details embedded inside of a registration.

```yaml
{"description": "Device details embedded inside of a registration.", "type": "object", "properties": {"client_version": {"description": "Version of the WARP client.", "type": "string", "example": "1.0.0", "x-auditable": true}, "id": {"description": "The ID of the device.", "type": "string", "example": "32aa0404-78f1-49a4-99e0-97f575081356", "x-auditable": true}, "name": {"description": "The name of the device.", "type": "string", "example": "My Device", "x-auditable": true}}, "required": ["id", "name"]}
```
