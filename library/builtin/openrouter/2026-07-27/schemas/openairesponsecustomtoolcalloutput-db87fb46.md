---
title: OpenAIResponseCustomToolCallOutput
page_id: schema-openairesponsecustomtoolcalloutput-db87fb46
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponseCustomToolCallOutput

```yaml
{"example": {"call_id": "call-abc123", "output": "patch applied successfully", "type": "custom_tool_call_output"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "output": {"anyOf": [{"type": "string"}, {"items": {"discriminator": {"mapping": {"input_file": "#/components/schemas/InputFile", "input_image": "#/components/schemas/InputImage", "input_text": "#/components/schemas/InputText"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/InputText"}, {"$ref": "#/components/schemas/InputImage"}, {"$ref": "#/components/schemas/InputFile"}]}, "type": "array"}]}, "type": {"enum": ["custom_tool_call_output"], "type": "string"}}, "required": ["type", "call_id", "output"], "type": "object"}
```
