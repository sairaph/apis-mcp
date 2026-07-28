---
title: WebSearchServerTool_OpenRouter
page_id: schema-websearchservertool-openrouter-a90eeb60
path: schemas
description: 'OpenRouter built-in server tool: searches the web for current information'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebSearchServerTool_OpenRouter

OpenRouter built-in server tool: searches the web for current information

```yaml
{"description": "OpenRouter built-in server tool: searches the web for current information", "example": {"parameters": {"max_results": 5}, "type": "openrouter:web_search"}, "properties": {"parameters": {"$ref": "#/components/schemas/WebSearchServerToolConfig"}, "type": {"enum": ["openrouter:web_search"], "type": "string"}}, "required": ["type"], "type": "object"}
```
