---
title: smartshield_smart_shield_settings_patch_body
page_id: schema-smartshield-smart-shield-settings-patch-body-ddc05600
path: schemas
description: The patch body for Smart Shield.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_smart_shield_settings_patch_body

The patch body for Smart Shield.

```yaml
{"description": "The patch body for Smart Shield.", "type": "object", "properties": {"cache_reserve": {"type": "object", "properties": {"value": {"description": "Specifies the enablement value of Cache Reserve.", "type": "string", "example": "on", "enum": ["on", "off"]}}}, "regional_tiered_cache": {"type": "object", "properties": {"value": {"description": "Specifies the enablement value of Regional Tiered Cache.", "type": "string", "example": "on", "enum": ["on", "off"]}}}, "smart_routing": {"type": "object", "properties": {"value": {"description": "Specifies the enablement value of Smart Routing.", "type": "string", "example": "on", "enum": ["on", "off"]}}}, "smart_tiered_cache": {"type": "object", "properties": {"value": {"description": "Specifies the enablement value of Smart Tiered Cache.", "type": "string", "example": "on", "enum": ["on", "off"]}}}}, "example": {"smart_tiered_cache": {"value": "on"}}}
```
