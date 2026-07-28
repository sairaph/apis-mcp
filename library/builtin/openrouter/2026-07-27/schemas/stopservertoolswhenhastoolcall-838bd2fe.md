---
title: StopServerToolsWhenHasToolCall
page_id: schema-stopservertoolswhenhastoolcall-838bd2fe
path: schemas
description: Stop after a tool with this name has been called.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenHasToolCall

Stop after a tool with this name has been called.

```yaml
{"description": "Stop after a tool with this name has been called.", "example": {"tool_name": "finalize", "type": "has_tool_call"}, "properties": {"tool_name": {"minLength": 1, "type": "string"}, "type": {"enum": ["has_tool_call"], "type": "string"}}, "required": ["type", "tool_name"], "type": "object"}
```
