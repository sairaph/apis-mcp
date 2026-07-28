---
title: ReasoningDetailText
page_id: schema-reasoningdetailtext-d113e712
path: schemas
description: Reasoning detail text schema
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDetailText

Reasoning detail text schema

```yaml
{"description": "Reasoning detail text schema", "example": {"signature": "signature", "text": "The model analyzed the problem by first identifying key constraints, then evaluating possible solutions...", "type": "reasoning.text"}, "properties": {"format": {"$ref": "#/components/schemas/ReasoningFormat"}, "id": {"type": ["string", "null"]}, "index": {"type": "integer"}, "signature": {"type": ["string", "null"]}, "text": {"type": ["string", "null"]}, "type": {"enum": ["reasoning.text"], "type": "string"}}, "required": ["type"], "type": "object"}
```
