---
title: smartshield_smart_shield_settings_get_response
page_id: schema-smartshield-smart-shield-settings-get-response-b0fa7889
path: schemas
description: The full Smart Shield response from the GET and PATCH operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_smart_shield_settings_get_response

The full Smart Shield response from the GET and PATCH operations.

```yaml
{"description": "The full Smart Shield response from the GET and PATCH operations.", "type": "object", "allOf": [{"$ref": "#/components/schemas/smartshield_smart_shield_settings"}, {"properties": {"healthchecks_count": {"description": "The total number of health checks associated with the zone.", "type": "integer", "example": 5}}, "type": "object"}], "required": ["smart_tiered_cache", "smart_routing", "regional_tiered_cache", "cache_reserve", "healthchecks_count"]}
```
