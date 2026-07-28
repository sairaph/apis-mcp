---
title: dlp_EmailRule
page_id: schema-dlp-emailrule-5d902ee0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EmailRule

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/dlp_EmailRuleAction"}, "conditions": {"description": "Triggered if all conditions match.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_EmailRuleCondition"}}, "created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string", "nullable": true}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "priority": {"type": "integer", "format": "int32", "minimum": 0}, "rule_id": {"type": "string", "format": "uuid"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["rule_id", "name", "priority", "enabled", "conditions", "action", "created_at", "updated_at"]}
```
