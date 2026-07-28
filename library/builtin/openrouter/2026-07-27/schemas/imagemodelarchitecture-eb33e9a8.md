---
title: ImageModelArchitecture
page_id: schema-imagemodelarchitecture-eb33e9a8
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageModelArchitecture

```yaml
{"example": {"input_modalities": ["text", "image"], "output_modalities": ["image"]}, "properties": {"input_modalities": {"description": "Supported input modalities", "items": {"$ref": "#/components/schemas/ImageInputModality"}, "type": "array"}, "output_modalities": {"description": "Supported output modalities", "items": {"$ref": "#/components/schemas/ImageOutputModality"}, "type": "array"}}, "required": ["input_modalities", "output_modalities"], "type": "object"}
```
