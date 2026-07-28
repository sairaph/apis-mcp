---
title: ObservabilityFilterRuleGroup
page_id: schema-observabilityfilterrulegroup-39550932
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ObservabilityFilterRuleGroup

```yaml
{"example": {"logic": "and", "rules": [{"field": "model", "operator": "equals", "value": "openai/gpt-4o"}]}, "properties": {"logic": {"default": "and", "enum": ["and", "or"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "rules": {"items": {"properties": {"field": {"enum": ["model", "provider", "session_id", "user_id", "api_key_name", "finish_reason", "input", "output", "total_cost", "total_tokens", "prompt_tokens", "completion_tokens"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "operator": {"enum": ["equals", "not_equals", "contains", "not_contains", "regex", "starts_with", "ends_with", "gt", "lt", "gte", "lte", "exists", "not_exists"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "value": {"anyOf": [{"type": "string"}, {"type": "number"}]}}, "required": ["field", "operator"], "type": "object"}, "type": "array"}}, "required": ["rules"], "type": "object"}
```
