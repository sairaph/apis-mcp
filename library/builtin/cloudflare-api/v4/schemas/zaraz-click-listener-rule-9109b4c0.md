---
title: zaraz_click-listener-rule
page_id: schema-zaraz-click-listener-rule-9109b4c0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_click-listener-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["clickListener"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"selector": {"type": "string", "x-auditable": true}, "type": {"type": "string", "enum": ["xpath", "css"], "x-auditable": true}, "waitForTags": {"type": "integer", "minimum": 0, "x-auditable": true}}, "required": ["type", "selector", "waitForTags"]}}, "required": ["id", "action", "settings"]}
```
