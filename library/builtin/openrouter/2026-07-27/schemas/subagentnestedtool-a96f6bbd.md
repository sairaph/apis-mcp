---
title: SubagentNestedTool
page_id: schema-subagentnestedtool-a96f6bbd
path: schemas
description: A tool made available to the subagent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the worker has no way to execute them. The subagent tool may not list itself.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SubagentNestedTool

A tool made available to the subagent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the worker has no way to execute them. The subagent tool may not list itself.

```yaml
{"additionalProperties": {}, "description": "A tool made available to the subagent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the worker has no way to execute them. The subagent tool may not list itself.", "example": {"type": "openrouter:web_search"}, "properties": {"parameters": {"additionalProperties": {}, "type": "object"}, "type": {"type": "string"}}, "required": ["type"], "type": "object"}
```
