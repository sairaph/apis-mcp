---
title: FusionServerTool_OpenRouter
page_id: schema-fusionservertool-openrouter-6e33e88d
path: schemas
description: 'OpenRouter built-in server tool: fans out the user prompt to a panel of analysis models, then asks a judge model to summarize their collective output as structured JSON the outer model can synthesize from.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionServerTool_OpenRouter

OpenRouter built-in server tool: fans out the user prompt to a panel of analysis models, then asks a judge model to summarize their collective output as structured JSON the outer model can synthesize from.

```yaml
{"description": "OpenRouter built-in server tool: fans out the user prompt to a panel of analysis models, then asks a judge model to summarize their collective output as structured JSON the outer model can synthesize from.", "example": {"parameters": {"analysis_models": ["~anthropic/claude-opus-latest", "~openai/gpt-latest"]}, "type": "openrouter:fusion"}, "properties": {"parameters": {"$ref": "#/components/schemas/FusionServerToolConfig"}, "type": {"enum": ["openrouter:fusion"], "type": "string"}}, "required": ["type"], "type": "object"}
```
