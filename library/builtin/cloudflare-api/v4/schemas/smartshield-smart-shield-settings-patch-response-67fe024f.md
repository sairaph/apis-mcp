---
title: smartshield_smart_shield_settings_patch_response
page_id: schema-smartshield-smart-shield-settings-patch-response-67fe024f
path: schemas
description: A consolidated object containing settings from multiple APIs for partial updates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_smart_shield_settings_patch_response

A consolidated object containing settings from multiple APIs for partial updates.

```yaml
{"description": "A consolidated object containing settings from multiple APIs for partial updates.", "type": "object", "properties": {"smart_tiered_cache": {"type": "object", "properties": {"editable": {"description": "Whether the setting is editable.", "type": "boolean", "example": true}, "id": {"description": "The id of the Smart Tiered Cache setting.", "type": "string", "example": "smart_tiered_cache"}, "modified_on": {"description": "The last time the setting was modified.", "type": "string", "example": "2025-09-10T22:53:22.946098Z"}, "value": {"description": "Specifies the enablement value of Tiered Cache.", "type": "string", "example": "on", "enum": ["on", "off"]}}}}, "required": ["smart_tiered_cache", "smart_routing", "regional_tiered_cache", "cache_reserve"]}
```
