---
title: FunctionTool
page_id: schema-functiontool-2e7e7eb1
path: schemas
description: Function tool definition
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FunctionTool

Function tool definition

```yaml
{"description": "Function tool definition", "example": {"description": "Get the current weather in a location", "name": "get_weather", "parameters": {"properties": {"location": {"description": "The city and state", "type": "string"}, "unit": {"enum": ["celsius", "fahrenheit"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["location"], "type": "object"}, "type": "function"}, "properties": {"description": {"type": ["string", "null"]}, "name": {"type": "string"}, "parameters": {"additionalProperties": {}, "type": ["object", "null"]}, "strict": {"type": ["boolean", "null"]}, "type": {"enum": ["function"], "type": "string"}}, "required": ["type", "name", "parameters"], "type": "object"}
```
