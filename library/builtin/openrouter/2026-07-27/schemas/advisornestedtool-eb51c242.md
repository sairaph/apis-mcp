---
title: AdvisorNestedTool
page_id: schema-advisornestedtool-eb51c242
path: schemas
description: A tool made available to the advisor sub-agent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the advisor has no way to execute them. The advisor tool may not list itself.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AdvisorNestedTool

A tool made available to the advisor sub-agent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the advisor has no way to execute them. The advisor tool may not list itself.

```yaml
{"additionalProperties": {}, "description": "A tool made available to the advisor sub-agent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the advisor has no way to execute them. The advisor tool may not list itself.", "example": {"type": "openrouter:web_search"}, "properties": {"parameters": {"additionalProperties": {}, "type": "object"}, "type": {"type": "string"}}, "required": ["type"], "type": "object"}
```
