---
title: cloudforce-one_RuleAction
page_id: schema-cloudforce-one-ruleaction-9a62f25d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_RuleAction

```yaml
{"type": "object", "properties": {"action_config": {"description": "Action-specific configuration parameters.", "type": "object", "additionalProperties": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}, {"items": {"type": "string"}, "type": "array"}, {"additionalProperties": {"type": "string"}, "type": "object"}]}}, "action_type": {"type": "string", "example": "alert_gchat", "enum": ["alert_gchat", "webhook", "logging", "email", "pipeline", "remediation", "throttle", "delete"]}, "enabled": {"type": "boolean", "default": true}}, "required": ["action_type", "action_config"]}
```
