---
title: ToolChoiceAllowed
page_id: schema-toolchoiceallowed-a345e245
path: schemas
description: Constrains the model to a pre-defined set of allowed tools
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ToolChoiceAllowed

Constrains the model to a pre-defined set of allowed tools

```yaml
{"description": "Constrains the model to a pre-defined set of allowed tools", "example": {"mode": "auto", "tools": [{"name": "get_weather", "type": "function"}], "type": "allowed_tools"}, "properties": {"mode": {"anyOf": [{"enum": ["auto"], "type": "string"}, {"enum": ["required"], "type": "string"}]}, "tools": {"items": {"additionalProperties": {}, "type": "object"}, "type": "array"}, "type": {"enum": ["allowed_tools"], "type": "string"}}, "required": ["type", "mode", "tools"], "type": "object"}
```
