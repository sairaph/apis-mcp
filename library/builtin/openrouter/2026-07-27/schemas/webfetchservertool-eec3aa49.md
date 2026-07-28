---
title: WebFetchServerTool
page_id: schema-webfetchservertool-eec3aa49
path: schemas
description: 'OpenRouter built-in server tool: fetches full content from a URL (web page or PDF)'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebFetchServerTool

OpenRouter built-in server tool: fetches full content from a URL (web page or PDF)

```yaml
{"description": "OpenRouter built-in server tool: fetches full content from a URL (web page or PDF)", "example": {"parameters": {"max_uses": 10}, "type": "openrouter:web_fetch"}, "properties": {"parameters": {"$ref": "#/components/schemas/WebFetchServerToolConfig"}, "type": {"enum": ["openrouter:web_fetch"], "type": "string"}}, "required": ["type"], "type": "object"}
```
