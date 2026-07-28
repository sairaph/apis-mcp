---
title: bot-management_bot_fight_mode_config
page_id: schema-bot-management-bot-fight-mode-config-84ac39e8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bot-management_bot_fight_mode_config

```yaml
{"allOf": [{"$ref": "#/components/schemas/bot-management_base_config"}, {"properties": {"fight_mode": {"$ref": "#/components/schemas/bot-management_fight_mode"}, "stale_zone_configuration": {"description": "A read-only field that shows which unauthorized settings are currently active on the zone. These settings typically result from upgrades or downgrades.", "type": "object", "properties": {"optimize_wordpress": {"$ref": "#/components/schemas/bot-management_optimize_wordpress_turned_on"}, "sbfm_definitely_automated": {"$ref": "#/components/schemas/bot-management_sbfm_definitely_automated_turned_on"}, "sbfm_likely_automated": {"$ref": "#/components/schemas/bot-management_sbfm_likely_automated_turned_on"}, "sbfm_static_resource_protection": {"$ref": "#/components/schemas/bot-management_sbfm_static_resource_protection_turned_on"}, "sbfm_verified_bots": {"$ref": "#/components/schemas/bot-management_sbfm_verified_bots_turned_on"}, "suppress_session_score": {"$ref": "#/components/schemas/bot-management_suppress_session_score_turned_off"}}, "readOnly": true, "title": "stale_zone_configuration"}}}], "title": "Bot Fight Mode"}
```
