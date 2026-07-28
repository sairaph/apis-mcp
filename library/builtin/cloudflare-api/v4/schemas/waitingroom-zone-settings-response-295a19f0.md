---
title: waitingroom_zone_settings_response
page_id: schema-waitingroom-zone-settings-response-295a19f0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_zone_settings_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"search_engine_crawler_bypass": {"$ref": "#/components/schemas/waitingroom_search_engine_crawler_bypass"}}, "required": ["search_engine_crawler_bypass"]}}, "required": ["result"], "type": "object"}]}
```
