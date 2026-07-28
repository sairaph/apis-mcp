---
title: SubagentServerToolConfig
page_id: schema-subagentservertoolconfig-e63d5884
path: schemas
description: Configuration for one openrouter:subagent server tool entry.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SubagentServerToolConfig

Configuration for one openrouter:subagent server tool entry.

```yaml
{"description": "Configuration for one openrouter:subagent server tool entry.", "example": {"model": "~anthropic/claude-haiku-latest", "name": "summarizer"}, "properties": {"instructions": {"description": "System instructions for the subagent. When omitted, the subagent responds with no system prompt of its own.", "example": "You are a fast, focused worker. Complete the task exactly as described.", "type": "string"}, "max_completion_tokens": {"description": "Maximum number of output tokens (including reasoning) the subagent may produce. When omitted, the provider's default applies.", "example": 2048, "type": "integer"}, "max_tool_calls": {"description": "Maximum number of tool-calling steps the subagent may take during its agentic loop. Capped at 25. Only relevant when the subagent is given tools. Accepted and validated but not yet enforced on the subagent call.", "example": 5, "maximum": 25, "minimum": 1, "type": "integer"}, "model": {"description": "Slug of the model that executes delegated tasks (any OpenRouter model). Typically a smaller, cheaper, faster model than the one delegating. When omitted, the model from the outer API request is used. The subagent tool itself cannot be the subagent model.", "example": "~anthropic/claude-haiku-latest", "type": "string"}, "name": {"description": "Optional name for this subagent. The model sees one tool per named subagent (and one default for an unnamed entry). Names must be unique across subagent entries. Letters, digits, spaces, underscores, and dashes; trimmed; 1–64 chars.", "example": "summarizer", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9 _-]+$", "type": "string"}, "reasoning": {"$ref": "#/components/schemas/SubagentReasoning"}, "temperature": {"description": "Sampling temperature forwarded to the subagent call. When omitted, the provider's default applies.", "example": 0.7, "format": "double", "type": "number"}, "tools": {"description": "Tools the subagent may use while executing a delegated task. The subagent runs as an agentic sub-agent over these tools, then returns its outcome. Only OpenRouter server tools are supported — function tools are rejected — and the list must not include the subagent tool itself.", "items": {"$ref": "#/components/schemas/SubagentNestedTool"}, "type": "array"}}, "type": "object"}
```
