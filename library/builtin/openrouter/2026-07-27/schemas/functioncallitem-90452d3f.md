---
title: FunctionCallItem
page_id: schema-functioncallitem-90452d3f
path: schemas
description: A function call initiated by the model
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FunctionCallItem

A function call initiated by the model

```yaml
{"allOf": [{"$ref": "#/components/schemas/OpenAIResponseFunctionToolCall"}, {"properties": {}, "required": ["id"], "type": "object"}], "description": "A function call initiated by the model", "example": {"arguments": "{\"location\":\"San Francisco\"}", "call_id": "call-abc123", "id": "call-abc123", "name": "get_weather", "status": "completed", "type": "function_call"}}
```
