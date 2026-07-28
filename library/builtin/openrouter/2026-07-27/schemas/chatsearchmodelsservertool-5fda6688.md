---
title: ChatSearchModelsServerTool
page_id: schema-chatsearchmodelsservertool-5fda6688
path: schemas
description: 'OpenRouter built-in server tool: searches and filters AI models available on OpenRouter'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatSearchModelsServerTool

OpenRouter built-in server tool: searches and filters AI models available on OpenRouter

```yaml
{"description": "OpenRouter built-in server tool: searches and filters AI models available on OpenRouter", "example": {"parameters": {"max_results": 5}, "type": "openrouter:experimental__search_models"}, "properties": {"parameters": {"$ref": "#/components/schemas/SearchModelsServerToolConfig"}, "type": {"enum": ["openrouter:experimental__search_models"], "type": "string"}}, "required": ["type"], "type": "object"}
```
