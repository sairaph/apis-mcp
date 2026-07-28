---
title: ModelReasoning
page_id: schema-modelreasoning-9eb86dba
path: schemas
description: Reasoning effort configuration. Omitted for non-reasoning models and dynamic router models.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ModelReasoning

Reasoning effort configuration. Omitted for non-reasoning models and dynamic router models.

```yaml
{"description": "Reasoning effort configuration. Omitted for non-reasoning models and dynamic router models.", "example": {"default_effort": "medium", "default_enabled": true, "mandatory": false, "supported_efforts": ["high", "medium", "low", "minimal"]}, "properties": {"default_effort": {"allOf": [{"$ref": "#/components/schemas/ReasoningEffort"}, {"description": "Default reasoning effort when the client enables reasoning without specifying effort. Maps to `reasoning.effort` in chat requests. When `\"none\"`, prefer omitting effort unless the user explicitly disables reasoning."}]}, "default_enabled": {"description": "Default reasoning enabled state when the client does not set `reasoning.enabled`.", "type": "boolean"}, "mandatory": {"description": "When true, reasoning cannot be disabled and effort \"none\" is rejected.", "type": "boolean"}, "supported_efforts": {"description": "Allowed reasoning effort values for this model, in descending effort order (highest first). Null means no allowlist — all gateway effort values are accepted.", "items": {"$ref": "#/components/schemas/ReasoningEffort"}, "type": ["array", "null"]}, "supports_max_tokens": {"description": "Present and `true` when the model accepts `reasoning.max_tokens` in requests (Anthropic-style) instead of or in addition to `reasoning.effort`. Omitted otherwise.", "type": "boolean"}}, "required": ["mandatory"], "type": "object"}
```
