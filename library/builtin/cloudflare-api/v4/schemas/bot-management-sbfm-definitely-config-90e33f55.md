---
title: bot-management_sbfm_definitely_config
page_id: schema-bot-management-sbfm-definitely-config-90e33f55
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bot-management_sbfm_definitely_config

```yaml
{"allOf": [{"$ref": "#/components/schemas/bot-management_base_config"}, {"properties": {"optimize_wordpress": {"$ref": "#/components/schemas/bot-management_optimize_wordpress"}, "sbfm_definitely_automated": {"$ref": "#/components/schemas/bot-management_sbfm_definitely_automated"}, "sbfm_static_resource_protection": {"$ref": "#/components/schemas/bot-management_sbfm_static_resource_protection"}, "sbfm_verified_bots": {"$ref": "#/components/schemas/bot-management_sbfm_verified_bots"}, "stale_zone_configuration": {"description": "A read-only field that shows which unauthorized settings are currently active on the zone. These settings typically result from upgrades or downgrades.", "type": "object", "properties": {"fight_mode": {"$ref": "#/components/schemas/bot-management_fight_mode_turned_on"}, "sbfm_likely_automated": {"$ref": "#/components/schemas/bot-management_sbfm_likely_automated_turned_on"}}, "readOnly": true, "title": "stale_zone_configuration"}}}], "title": "SBFM Pro Plan"}
```
