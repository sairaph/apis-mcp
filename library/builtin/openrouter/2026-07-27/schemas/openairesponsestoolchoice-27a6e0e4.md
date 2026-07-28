---
title: OpenAIResponsesToolChoice
page_id: schema-openairesponsestoolchoice-27a6e0e4
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponsesToolChoice

```yaml
{"anyOf": [{"enum": ["auto"], "type": "string"}, {"enum": ["none"], "type": "string"}, {"enum": ["required"], "type": "string"}, {"properties": {"name": {"type": "string"}, "type": {"enum": ["function"], "type": "string"}}, "required": ["type", "name"], "type": "object"}, {"properties": {"type": {"anyOf": [{"enum": ["web_search_preview_2025_03_11"], "type": "string"}, {"enum": ["web_search_preview"], "type": "string"}]}}, "required": ["type"], "type": "object"}, {"$ref": "#/components/schemas/ToolChoiceAllowed"}, {"properties": {"type": {"enum": ["apply_patch"], "type": "string"}}, "required": ["type"], "type": "object"}, {"properties": {"type": {"enum": ["shell"], "type": "string"}}, "required": ["type"], "type": "object"}], "example": "auto"}
```
