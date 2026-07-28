---
title: GenerationContentResponse
page_id: schema-generationcontentresponse-cb82fe12
path: schemas
description: Stored prompt and completion content for a generation
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GenerationContentResponse

Stored prompt and completion content for a generation

```yaml
{"description": "Stored prompt and completion content for a generation", "example": {"data": {"input": {"messages": [{"content": "What is the meaning of life?", "role": "user"}]}, "output": {"completion": "The meaning of life is a philosophical question...", "reasoning": null}}}, "properties": {"data": {"$ref": "#/components/schemas/GenerationContentData"}}, "required": ["data"], "type": "object"}
```
