---
title: OpenAIResponseFunctionToolCallOutput
page_id: schema-openairesponsefunctiontoolcalloutput-7e88bba4
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponseFunctionToolCallOutput

```yaml
{"example": {"call_id": "call-abc123", "output": "{\"temperature\":72,\"conditions\":\"sunny\"}", "type": "function_call_output"}, "properties": {"call_id": {"type": "string"}, "id": {"type": ["string", "null"]}, "output": {"anyOf": [{"type": "string"}, {"items": {"discriminator": {"mapping": {"input_file": "#/components/schemas/InputFile", "input_image": "#/components/schemas/InputImage", "input_text": "#/components/schemas/InputText"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/InputText"}, {"$ref": "#/components/schemas/InputImage"}, {"$ref": "#/components/schemas/InputFile"}]}, "type": "array"}]}, "status": {"anyOf": [{"$ref": "#/components/schemas/ToolCallStatus"}, {"type": "null"}]}, "type": {"enum": ["function_call_output"], "type": "string"}}, "required": ["type", "call_id", "output"], "type": "object"}
```
