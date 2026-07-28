---
title: MessagesToolRemovalBlock
page_id: schema-messagestoolremovalblock-a85b59fb
path: schemas
description: 'Removes a tool from the conversation mid-conversation without invalidating the prompt cache. Only valid in `role: "system"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesToolRemovalBlock

Removes a tool from the conversation mid-conversation without invalidating the prompt cache. Only valid in `role: "system"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.

```yaml
{"description": "Removes a tool from the conversation mid-conversation without invalidating the prompt cache. Only valid in `role: \"system\"` messages. Not supported on Claude Sonnet 5 or models older than Claude Opus 4.8.", "example": {"tool": {"name": "get_weather", "type": "tool_reference"}, "type": "tool_removal"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "tool": {"oneOf": [{"properties": {"name": {"type": "string"}, "type": {"enum": ["tool_reference"], "type": "string"}}, "required": ["type", "name"], "type": "object"}, {"properties": {"name": {"type": "string"}, "server_name": {"type": "string"}, "type": {"enum": ["mcp_tool_reference"], "type": "string"}}, "required": ["type", "name", "server_name"], "type": "object"}, {"properties": {"server_name": {"type": "string"}, "type": {"enum": ["mcp_toolset_reference"], "type": "string"}}, "required": ["type", "server_name"], "type": "object"}]}, "type": {"enum": ["tool_removal"], "type": "string"}}, "required": ["type", "tool"], "type": "object"}
```
