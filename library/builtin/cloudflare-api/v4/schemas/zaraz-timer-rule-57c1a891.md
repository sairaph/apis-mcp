---
title: zaraz_timer-rule
page_id: schema-zaraz-timer-rule-57c1a891
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_timer-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["timer"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"interval": {"type": "integer", "minimum": 50, "x-auditable": true}, "limit": {"type": "integer", "minimum": 0, "x-auditable": true}}, "required": ["interval", "limit"]}}, "required": ["id", "action", "settings"]}
```
