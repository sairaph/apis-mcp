---
title: zaraz_scroll-depth-rule
page_id: schema-zaraz-scroll-depth-rule-a8c25fd6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_scroll-depth-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["scrollDepth"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"positions": {"type": "string", "x-auditable": true}}, "required": ["positions"]}}, "required": ["id", "action", "settings"]}
```
