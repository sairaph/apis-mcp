---
title: FilesServerTool
page_id: schema-filesservertool-f0d3e3fb
path: schemas
description: 'OpenRouter built-in server tool: read, write, edit, and list workspace files via the Files API. Requires the `x-openrouter-file-ids: openrouter` request header.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FilesServerTool

OpenRouter built-in server tool: read, write, edit, and list workspace files via the Files API. Requires the `x-openrouter-file-ids: openrouter` request header.

```yaml
{"description": "OpenRouter built-in server tool: read, write, edit, and list workspace files via the Files API. Requires the `x-openrouter-file-ids: openrouter` request header.", "example": {"parameters": {}, "type": "openrouter:files"}, "properties": {"parameters": {"$ref": "#/components/schemas/FilesServerToolConfig"}, "type": {"enum": ["openrouter:files"], "type": "string"}}, "required": ["type"], "type": "object"}
```
