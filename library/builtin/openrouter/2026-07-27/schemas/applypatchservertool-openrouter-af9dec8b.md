---
title: ApplyPatchServerTool_OpenRouter
page_id: schema-applypatchservertool-openrouter-af9dec8b
path: schemas
description: 'OpenRouter built-in server tool: validates V4A diff patches for file operations (create, update, delete). Restricted to the Responses API.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchServerTool_OpenRouter

OpenRouter built-in server tool: validates V4A diff patches for file operations (create, update, delete). Restricted to the Responses API.

```yaml
{"description": "OpenRouter built-in server tool: validates V4A diff patches for file operations (create, update, delete). Restricted to the Responses API.", "example": {"type": "openrouter:apply_patch"}, "properties": {"parameters": {"$ref": "#/components/schemas/ApplyPatchServerToolConfig"}, "type": {"enum": ["openrouter:apply_patch"], "type": "string"}}, "required": ["type"], "type": "object"}
```
