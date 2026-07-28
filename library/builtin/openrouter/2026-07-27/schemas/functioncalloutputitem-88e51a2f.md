---
title: FunctionCallOutputItem
page_id: schema-functioncalloutputitem-88e51a2f
path: schemas
description: The output from a function call execution
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FunctionCallOutputItem

The output from a function call execution

```yaml
{"allOf": [{"$ref": "#/components/schemas/OpenAIResponseFunctionToolCallOutput"}, {"properties": {"output": {"anyOf": [{"type": "string"}, {"items": {"oneOf": [{"$ref": "#/components/schemas/InputText"}, {"allOf": [{"$ref": "#/components/schemas/InputImage"}, {"properties": {}, "type": "object"}], "description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}}, {"$ref": "#/components/schemas/InputFile"}]}, "type": "array"}]}}, "type": "object"}], "description": "The output from a function call execution", "example": {"call_id": "call-abc123", "id": "output-abc123", "output": "{\"temperature\":72,\"conditions\":\"sunny\"}", "status": "completed", "type": "function_call_output"}}
```
