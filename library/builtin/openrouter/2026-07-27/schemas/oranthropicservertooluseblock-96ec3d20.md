---
title: ORAnthropicServerToolUseBlock
page_id: schema-oranthropicservertooluseblock-96ec3d20
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ORAnthropicServerToolUseBlock

```yaml
{"example": {"caller": {"type": "direct"}, "id": "srvtoolu_01abc", "input": {}, "name": "advisor", "type": "server_tool_use"}, "properties": {"caller": {"$ref": "#/components/schemas/ORAnthropicNullableCaller"}, "id": {"type": "string"}, "input": {}, "name": {"type": "string"}, "type": {"enum": ["server_tool_use"], "type": "string"}}, "required": ["type", "id", "name"], "type": "object"}
```
