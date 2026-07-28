---
title: BaseInputs
page_id: schema-baseinputs-59d5d1b7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseInputs

```yaml
{"anyOf": [{"type": "string"}, {"items": {"anyOf": [{"properties": {"content": {"anyOf": [{"items": {"discriminator": {"mapping": {"input_audio": "#/components/schemas/InputAudio", "input_file": "#/components/schemas/InputFile", "input_image": "#/components/schemas/InputImage", "input_text": "#/components/schemas/InputText"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/InputText"}, {"$ref": "#/components/schemas/InputImage"}, {"$ref": "#/components/schemas/InputFile"}, {"$ref": "#/components/schemas/InputAudio"}]}, "type": "array"}, {"type": "string"}]}, "phase": {"anyOf": [{"enum": ["commentary"], "type": "string"}, {"enum": ["final_answer"], "type": "string"}, {"type": "null"}]}, "role": {"anyOf": [{"enum": ["user"], "type": "string"}, {"enum": ["system"], "type": "string"}, {"enum": ["assistant"], "type": "string"}, {"enum": ["developer"], "type": "string"}]}, "type": {"enum": ["message"], "type": "string"}}, "required": ["role", "content"], "type": "object"}, {"$ref": "#/components/schemas/OpenAIResponseInputMessageItem"}, {"$ref": "#/components/schemas/OpenAIResponseFunctionToolCallOutput"}, {"$ref": "#/components/schemas/OpenAIResponseFunctionToolCall"}, {"$ref": "#/components/schemas/OutputItemImageGenerationCall"}, {"$ref": "#/components/schemas/OutputMessage"}, {"$ref": "#/components/schemas/OpenAIResponseCustomToolCall"}, {"$ref": "#/components/schemas/OpenAIResponseCustomToolCallOutput"}, {"$ref": "#/components/schemas/ApplyPatchCallItem"}, {"$ref": "#/components/schemas/ApplyPatchCallOutputItem"}]}, "type": "array"}, {"type": "null"}], "example": [{"content": "What is the weather today?", "role": "user"}]}
```
