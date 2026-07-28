---
title: ChatServerToolChoice
page_id: schema-chatservertoolchoice-69b2b71c
path: schemas
description: 'OpenRouter extension: force a specific server tool by naming it directly in `tool_choice.type` instead of wrapping it in `{ type: "function", function: { name } }`.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatServerToolChoice

OpenRouter extension: force a specific server tool by naming it directly in `tool_choice.type` instead of wrapping it in `{ type: "function", function: { name } }`.

```yaml
{"description": "OpenRouter extension: force a specific server tool by naming it directly in `tool_choice.type` instead of wrapping it in `{ type: \"function\", function: { name } }`.", "example": {"type": "openrouter:web_search"}, "properties": {"type": {"description": "OpenRouter server-tool type to force (e.g. `openrouter:web_search`, `web_search`, `web_search_preview`).", "example": "openrouter:web_search", "type": "string"}}, "required": ["type"], "type": "object"}
```
