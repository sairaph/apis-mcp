---
title: ShellServerTool_OpenRouter
page_id: schema-shellservertool-openrouter-53b46285
path: schemas
description: 'OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container (a sandbox-backed clone of OpenAI''s hosted shell tool)'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellServerTool_OpenRouter

OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container (a sandbox-backed clone of OpenAI's hosted shell tool)

```yaml
{"description": "OpenRouter built-in server tool: runs shell commands server-side in a sandboxed container (a sandbox-backed clone of OpenAI's hosted shell tool)", "example": {"parameters": {"engine": "openrouter", "environment": {"type": "container_auto"}}, "type": "openrouter:shell"}, "properties": {"parameters": {"$ref": "#/components/schemas/ShellServerToolConfig"}, "type": {"enum": ["openrouter:shell"], "type": "string"}}, "required": ["type"], "type": "object"}
```
