---
title: SubagentServerTool_OpenRouter
page_id: schema-subagentservertool-openrouter-f93bf274
path: schemas
description: 'OpenRouter built-in server tool: delegates self-contained tasks to a smaller, cheaper, faster worker model (any OpenRouter model) mid-generation and returns its outcome. The worker may run as a sub-agent with its own tools.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SubagentServerTool_OpenRouter

OpenRouter built-in server tool: delegates self-contained tasks to a smaller, cheaper, faster worker model (any OpenRouter model) mid-generation and returns its outcome. The worker may run as a sub-agent with its own tools.

```yaml
{"description": "OpenRouter built-in server tool: delegates self-contained tasks to a smaller, cheaper, faster worker model (any OpenRouter model) mid-generation and returns its outcome. The worker may run as a sub-agent with its own tools.", "example": {"parameters": {"model": "~anthropic/claude-haiku-latest"}, "type": "openrouter:subagent"}, "properties": {"parameters": {"$ref": "#/components/schemas/SubagentServerToolConfig"}, "type": {"enum": ["openrouter:subagent"], "type": "string"}}, "required": ["type"], "type": "object"}
```
