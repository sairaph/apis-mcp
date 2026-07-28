---
title: teams-devices_virtual_networks
page_id: schema-teams-devices-virtual-networks-65ef7373
path: schemas
description: Virtual network access settings for the device.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_virtual_networks

Virtual network access settings for the device.

```yaml
{"description": "Virtual network access settings for the device.", "type": "object", "properties": {"allowed": {"description": "List of virtual network IDs the device is allowed to access. When virtual_networks is set, at least one entry is required.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"], "minItems": 1, "uniqueItems": true}, "default": {"description": "The default virtual network ID. Must be included in the `allowed` list.", "type": "string", "format": "uuid", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}, "nullable": true, "required": ["allowed", "default"]}
```
