---
title: OpenRouterWebSearchServerTool
page_id: schema-openrouterwebsearchservertool-f1443b69
path: schemas
description: 'OpenRouter built-in server tool: searches the web for current information'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenRouterWebSearchServerTool

OpenRouter built-in server tool: searches the web for current information

```yaml
{"description": "OpenRouter built-in server tool: searches the web for current information", "example": {"parameters": {"max_results": 5}, "type": "openrouter:web_search"}, "properties": {"parameters": {"$ref": "#/components/schemas/WebSearchConfig"}, "type": {"enum": ["openrouter:web_search"], "type": "string"}}, "required": ["type"], "type": "object"}
```
