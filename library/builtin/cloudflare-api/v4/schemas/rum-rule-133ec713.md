---
title: rum_rule
page_id: schema-rum-rule-133ec713
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_rule

```yaml
{"type": "object", "properties": {"created": {"$ref": "#/components/schemas/rum_timestamp"}, "host": {"description": "The hostname the rule will be applied to.", "type": "string", "example": "example.com", "x-auditable": true}, "id": {"$ref": "#/components/schemas/rum_rule_identifier"}, "inclusive": {"description": "Whether the rule includes or excludes traffic from being measured.", "type": "boolean", "example": true, "x-auditable": true}, "is_paused": {"description": "Whether the rule is paused or not.", "type": "boolean", "example": false, "x-auditable": true}, "paths": {"description": "The paths the rule will be applied to.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["*"]}, "priority": {"type": "number", "example": 1000, "x-auditable": true}}}
```
