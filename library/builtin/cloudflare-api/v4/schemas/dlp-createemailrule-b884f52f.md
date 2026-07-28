---
title: dlp_CreateEmailRule
page_id: schema-dlp-createemailrule-b884f52f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CreateEmailRule

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/dlp_EmailRuleAction"}, "conditions": {"description": "Triggered if all conditions match.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_EmailRuleCondition"}}, "description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}}, "required": ["name", "enabled", "conditions", "action"]}
```
