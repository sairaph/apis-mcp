---
title: GenerationContentData
page_id: schema-generationcontentdata-f72a4168
path: schemas
description: Stored prompt and completion content
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GenerationContentData

Stored prompt and completion content

```yaml
{"description": "Stored prompt and completion content", "example": {"input": {"messages": [{"content": "What is the meaning of life?", "role": "user"}]}, "output": {"completion": "The meaning of life is a philosophical question...", "reasoning": null}}, "properties": {"input": {"anyOf": [{"properties": {"prompt": {"example": "What is the meaning of life?", "type": "string"}}, "required": ["prompt"], "type": "object"}, {"properties": {"messages": {"example": [{"content": "What is the meaning of life?", "role": "user"}], "items": {}, "type": "array"}}, "required": ["messages"], "type": "object"}], "description": "The input to the generation — either a prompt string or an array of messages"}, "output": {"description": "The output from the generation", "properties": {"completion": {"description": "The completion output", "example": "The meaning of life is a philosophical question...", "type": ["string", "null"]}, "reasoning": {"description": "Reasoning/thinking output, if any", "example": null, "type": ["string", "null"]}}, "required": ["reasoning", "completion"], "type": "object"}}, "required": ["input", "output"], "type": "object"}
```
