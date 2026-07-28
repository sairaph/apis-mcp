---
title: AnthropicToolUseBlock
page_id: schema-anthropictooluseblock-9d8f7142
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolUseBlock

```yaml
{"example": {"caller": {"type": "direct"}, "id": "toolu_01abc", "input": {"location": "San Francisco"}, "name": "get_weather", "type": "tool_use"}, "properties": {"caller": {"$ref": "#/components/schemas/AnthropicCaller"}, "id": {"type": "string"}, "input": {}, "name": {"type": "string"}, "type": {"enum": ["tool_use"], "type": "string"}}, "required": ["type", "id", "caller", "name"], "type": "object"}
```
