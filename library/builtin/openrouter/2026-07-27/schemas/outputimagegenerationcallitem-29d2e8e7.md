---
title: OutputImageGenerationCallItem
page_id: schema-outputimagegenerationcallitem-29d2e8e7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputImageGenerationCallItem

```yaml
{"allOf": [{"$ref": "#/components/schemas/OutputItemImageGenerationCall"}, {"properties": {"prompt": {"description": "The prompt (possibly rewritten) that the image was generated from.", "type": "string"}}, "type": "object"}], "example": {"id": "img-abc123", "result": null, "status": "completed", "type": "image_generation_call"}}
```
