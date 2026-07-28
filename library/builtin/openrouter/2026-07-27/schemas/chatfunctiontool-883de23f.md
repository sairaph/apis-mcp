---
title: ChatFunctionTool
page_id: schema-chatfunctiontool-883de23f
path: schemas
description: Tool definition for function calling (regular function or OpenRouter built-in server tool)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatFunctionTool

Tool definition for function calling (regular function or OpenRouter built-in server tool)

```yaml
{"anyOf": [{"properties": {"cache_control": {"$ref": "#/components/schemas/ChatContentCacheControl"}, "function": {"description": "Function definition for tool calling", "example": {"description": "Get the current weather for a location", "name": "get_weather", "parameters": {"properties": {"location": {"description": "City name", "type": "string"}}, "required": ["location"], "type": "object"}}, "properties": {"description": {"description": "Function description for the model", "example": "Get the current weather for a location", "type": "string"}, "name": {"description": "Function name (a-z, A-Z, 0-9, underscores, dashes, max 64 chars)", "example": "get_weather", "maxLength": 64, "type": "string"}, "parameters": {"additionalProperties": {}, "description": "Function parameters as JSON Schema object", "example": {"properties": {"location": {"description": "City name", "type": "string"}}, "required": ["location"], "type": "object"}, "type": "object"}, "strict": {"description": "Enable strict schema adherence", "example": false, "type": ["boolean", "null"]}}, "required": ["name"], "type": "object"}, "type": {"enum": ["function"], "type": "string"}}, "required": ["type", "function"], "type": "object"}, {"$ref": "#/components/schemas/AdvisorServerTool_OpenRouter"}, {"$ref": "#/components/schemas/BashServerTool"}, {"$ref": "#/components/schemas/DatetimeServerTool"}, {"$ref": "#/components/schemas/FilesServerTool"}, {"$ref": "#/components/schemas/FusionServerTool_OpenRouter"}, {"$ref": "#/components/schemas/ImageGenerationServerTool_OpenRouter"}, {"$ref": "#/components/schemas/ChatSearchModelsServerTool"}, {"$ref": "#/components/schemas/SubagentServerTool_OpenRouter"}, {"$ref": "#/components/schemas/WebFetchServerTool"}, {"$ref": "#/components/schemas/OpenRouterWebSearchServerTool"}, {"$ref": "#/components/schemas/ChatWebSearchShorthand"}], "description": "Tool definition for function calling (regular function or OpenRouter built-in server tool)", "example": {"function": {"description": "Get the current weather for a location", "name": "get_weather", "parameters": {"properties": {"location": {"description": "City name", "type": "string"}, "unit": {"enum": ["celsius", "fahrenheit"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["location"], "type": "object"}}, "type": "function"}}
```
