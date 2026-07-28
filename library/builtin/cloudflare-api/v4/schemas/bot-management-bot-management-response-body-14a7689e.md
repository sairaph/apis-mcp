---
title: bot-management_bot_management_response_body
page_id: schema-bot-management-bot-management-response-body-14a7689e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bot-management_bot_management_response_body

```yaml
{"allOf": [{"$ref": "#/components/schemas/bot-management_api-response-single"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/bot-management_bot_fight_mode_config"}, {"$ref": "#/components/schemas/bot-management_sbfm_definitely_config"}, {"$ref": "#/components/schemas/bot-management_sbfm_likely_config"}, {"$ref": "#/components/schemas/bot-management_bm_subscription_config"}]}}, "type": "object"}]}
```
