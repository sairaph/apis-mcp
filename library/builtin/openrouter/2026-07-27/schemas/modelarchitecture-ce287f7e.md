---
title: ModelArchitecture
page_id: schema-modelarchitecture-ce287f7e
path: schemas
description: Model architecture information
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ModelArchitecture

Model architecture information

```yaml
{"description": "Model architecture information", "example": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "properties": {"input_modalities": {"description": "Supported input modalities", "items": {"$ref": "#/components/schemas/InputModality"}, "type": "array"}, "instruct_type": {"$ref": "#/components/schemas/InstructType"}, "modality": {"description": "Primary modality of the model", "example": "text->text", "type": ["string", "null"]}, "output_modalities": {"description": "Supported output modalities", "items": {"$ref": "#/components/schemas/OutputModality"}, "type": "array"}, "tokenizer": {"$ref": "#/components/schemas/ModelGroup"}}, "required": ["modality", "input_modalities", "output_modalities"], "type": "object"}
```
