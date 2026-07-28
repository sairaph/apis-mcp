---
title: BashServerTool
page_id: schema-bashservertool-2d1080ef
path: schemas
description: 'OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BashServerTool

OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container

```yaml
{"description": "OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container", "example": {"parameters": {"environment": {"type": "container_auto"}}, "type": "openrouter:bash"}, "properties": {"parameters": {"$ref": "#/components/schemas/BashServerToolConfig"}, "type": {"enum": ["openrouter:bash"], "type": "string"}}, "required": ["type"], "type": "object"}
```
