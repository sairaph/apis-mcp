---
title: ReasoningItem
page_id: schema-reasoningitem-fd798e23
path: schemas
description: Reasoning output item with signature and format extensions
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningItem

Reasoning output item with signature and format extensions

```yaml
{"allOf": [{"$ref": "#/components/schemas/OutputItemReasoning"}, {"properties": {"content": {"items": {"$ref": "#/components/schemas/ReasoningTextContent"}, "type": ["array", "null"]}, "format": {"$ref": "#/components/schemas/ReasoningFormat"}, "signature": {"type": ["string", "null"]}}, "type": "object"}], "description": "Reasoning output item with signature and format extensions", "example": {"id": "reasoning-abc123", "summary": [{"text": "Step by step analysis", "type": "summary_text"}], "type": "reasoning"}}
```
