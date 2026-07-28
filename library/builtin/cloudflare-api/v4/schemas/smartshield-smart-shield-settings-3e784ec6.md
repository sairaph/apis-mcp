---
title: smartshield_smart_shield_settings
page_id: schema-smartshield-smart-shield-settings-3e784ec6
path: schemas
description: A consolidated object containing settings from multiple APIs for partial updates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_smart_shield_settings

A consolidated object containing settings from multiple APIs for partial updates.

```yaml
{"description": "A consolidated object containing settings from multiple APIs for partial updates.", "type": "object", "properties": {"cache_reserve": {"properties": {"editable": {"description": "Whether the setting is editable.", "type": "boolean", "example": true}, "id": {"description": "The id of the Cache Reserve setting.", "type": "string", "example": "cache_reserve"}, "value": {"description": "Specifies the enablement value of Cache Reserve.", "type": "string", "example": "off", "enum": ["on", "off"]}}}, "regional_tiered_cache": {"type": "object", "properties": {"editable": {"description": "Whether the setting is editable.", "type": "boolean", "example": true}, "id": {"description": "The id of the Regional Tiered Cache setting.", "type": "string", "example": "regional_tiered_cache"}, "value": {"description": "Specifies the enablement value of Cache Reserve.", "type": "string", "example": "off", "enum": ["on", "off"]}}}, "smart_routing": {"type": "object", "properties": {"editable": {"description": "Whether the setting is editable.", "type": "boolean", "example": true}, "id": {"description": "The id of the Smart Routing setting.", "type": "string", "example": "smart_routing"}, "value": {"description": "Specifies the enablement value of Argo Smart Routing.", "type": "string", "example": "off", "enum": ["on", "off"]}}}, "smart_tiered_cache": {"type": "object", "properties": {"editable": {"description": "Whether the setting is editable.", "type": "boolean", "example": true}, "id": {"description": "The id of the Smart Tiered Cache setting.", "type": "string", "example": "smart_tiered_cache"}, "modified_on": {"description": "The last time the setting was modified.", "type": "string", "example": "2025-09-10T22:53:22.946098Z"}, "value": {"description": "Specifies the enablement value of Tiered Cache.", "type": "string", "example": "on", "enum": ["on", "off"]}}}}}
```
