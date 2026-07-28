---
title: ChatNamedToolChoice
page_id: schema-chatnamedtoolchoice-e930e3cc
path: schemas
description: Named tool choice for specific function
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatNamedToolChoice

Named tool choice for specific function

```yaml
{"description": "Named tool choice for specific function", "example": {"function": {"name": "get_weather"}, "type": "function"}, "properties": {"function": {"properties": {"name": {"description": "Function name to call", "example": "get_weather", "type": "string"}}, "required": ["name"], "type": "object"}, "type": {"enum": ["function"], "type": "string"}}, "required": ["type", "function"], "type": "object"}
```
