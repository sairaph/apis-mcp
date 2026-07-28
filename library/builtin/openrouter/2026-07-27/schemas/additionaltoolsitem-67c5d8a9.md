---
title: AdditionalToolsItem
page_id: schema-additionaltoolsitem-67c5d8a9
path: schemas
description: Additional tools made available to the model at this point in the input
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AdditionalToolsItem

Additional tools made available to the model at this point in the input

```yaml
{"description": "Additional tools made available to the model at this point in the input", "example": {"role": "developer", "tools": [{"name": "get_weather", "parameters": {"properties": {"location": {"type": "string"}}, "type": "object"}, "type": "function"}], "type": "additional_tools"}, "properties": {"id": {"type": ["string", "null"]}, "role": {"enum": ["unknown", "user", "assistant", "system", "critic", "discriminator", "developer", "tool"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "tools": {"items": {"anyOf": [{"allOf": [{"$ref": "#/components/schemas/FunctionTool"}, {"properties": {}, "type": "object"}], "description": "Function tool definition", "example": {"description": "Get the current weather in a location", "name": "get_weather", "parameters": {"properties": {"location": {"description": "The city and state", "type": "string"}, "unit": {"enum": ["celsius", "fahrenheit"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["location"], "type": "object"}, "type": "function"}}, {"$ref": "#/components/schemas/Preview_WebSearchServerTool"}, {"$ref": "#/components/schemas/Preview_20250311_WebSearchServerTool"}, {"$ref": "#/components/schemas/Legacy_WebSearchServerTool"}, {"$ref": "#/components/schemas/WebSearchServerTool"}, {"$ref": "#/components/schemas/FileSearchServerTool"}, {"$ref": "#/components/schemas/ComputerUseServerTool"}, {"$ref": "#/components/schemas/CodeInterpreterServerTool"}, {"$ref": "#/components/schemas/McpServerTool"}, {"$ref": "#/components/schemas/ImageGenerationServerTool"}, {"$ref": "#/components/schemas/CodexLocalShellTool"}, {"$ref": "#/components/schemas/ShellServerTool"}, {"$ref": "#/components/schemas/ApplyPatchServerTool"}, {"$ref": "#/components/schemas/CustomTool"}, {"$ref": "#/components/schemas/NamespaceTool"}, {"$ref": "#/components/schemas/AdvisorServerTool_OpenRouter"}, {"$ref": "#/components/schemas/SubagentServerTool_OpenRouter"}, {"$ref": "#/components/schemas/DatetimeServerTool"}, {"$ref": "#/components/schemas/FilesServerTool"}, {"$ref": "#/components/schemas/FusionServerTool_OpenRouter"}, {"$ref": "#/components/schemas/ImageGenerationServerTool_OpenRouter"}, {"$ref": "#/components/schemas/SearchModelsServerTool_OpenRouter"}, {"$ref": "#/components/schemas/WebFetchServerTool"}, {"$ref": "#/components/schemas/WebSearchServerTool_OpenRouter"}, {"$ref": "#/components/schemas/ApplyPatchServerTool_OpenRouter"}, {"$ref": "#/components/schemas/BashServerTool"}, {"$ref": "#/components/schemas/ShellServerTool_OpenRouter"}, {"additionalProperties": {}, "properties": {"type": {"type": "string"}}, "required": ["type"], "type": "object"}]}, "type": "array"}, "type": {"enum": ["additional_tools"], "type": "string"}}, "required": ["type", "role", "tools"], "type": "object"}
```
