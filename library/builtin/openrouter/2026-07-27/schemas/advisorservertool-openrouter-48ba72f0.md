---
title: AdvisorServerTool_OpenRouter
page_id: schema-advisorservertool-openrouter-48ba72f0
path: schemas
description: 'OpenRouter built-in server tool: consults a higher-intelligence advisor model (any OpenRouter model) for guidance mid-generation and returns its response. The advisor may run as a sub-agent with its own tools. Include multiple entries to offer several named advisors; at most one entry may omit `name` to act as the default advisor.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AdvisorServerTool_OpenRouter

OpenRouter built-in server tool: consults a higher-intelligence advisor model (any OpenRouter model) for guidance mid-generation and returns its response. The advisor may run as a sub-agent with its own tools. Include multiple entries to offer several named advisors; at most one entry may omit `name` to act as the default advisor.

```yaml
{"description": "OpenRouter built-in server tool: consults a higher-intelligence advisor model (any OpenRouter model) for guidance mid-generation and returns its response. The advisor may run as a sub-agent with its own tools. Include multiple entries to offer several named advisors; at most one entry may omit `name` to act as the default advisor.", "example": {"parameters": {"model": "~anthropic/claude-opus-latest", "name": "reviewer"}, "type": "openrouter:advisor"}, "properties": {"parameters": {"$ref": "#/components/schemas/AdvisorServerToolConfig"}, "type": {"enum": ["openrouter:advisor"], "type": "string"}}, "required": ["type"], "type": "object"}
```
