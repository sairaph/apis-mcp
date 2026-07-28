---
title: waf-managed-rules_schemas-group
page_id: schema-waf-managed-rules-schemas-group-d85b8433
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_schemas-group

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/waf-managed-rules_group"}, {"properties": {"allowed_modes": {"$ref": "#/components/schemas/waf-managed-rules_allowed_modes"}, "mode": {"$ref": "#/components/schemas/waf-managed-rules_mode"}}, "type": "object"}], "required": ["id", "name", "description", "mode", "rules_count"]}
```
