---
title: MessagesToolAdditionBlock
page_id: schema-messagestooladditionblock-84fa597d
path: schemas
description: 'Loads a previously deferred tool (declared in `tools` with `defer_loading: true`) mid-conversation without invalidating the prompt cache. Only valid in `role: "system"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesToolAdditionBlock

Loads a previously deferred tool (declared in `tools` with `defer_loading: true`) mid-conversation without invalidating the prompt cache. Only valid in `role: "system"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.

```yaml
{"description": "Loads a previously deferred tool (declared in `tools` with `defer_loading: true`) mid-conversation without invalidating the prompt cache. Only valid in `role: \"system\"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.", "example": {"tool": {"name": "get_forecast", "type": "tool_reference"}, "type": "tool_addition"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "tool": {"oneOf": [{"properties": {"name": {"type": "string"}, "type": {"enum": ["tool_reference"], "type": "string"}}, "required": ["type", "name"], "type": "object"}, {"properties": {"name": {"type": "string"}, "server_name": {"type": "string"}, "type": {"enum": ["mcp_tool_reference"], "type": "string"}}, "required": ["type", "name", "server_name"], "type": "object"}, {"properties": {"server_name": {"type": "string"}, "type": {"enum": ["mcp_toolset_reference"], "type": "string"}}, "required": ["type", "server_name"], "type": "object"}]}, "type": {"enum": ["tool_addition"], "type": "string"}}, "required": ["type", "tool"], "type": "object"}
```
