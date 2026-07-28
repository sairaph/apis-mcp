---
title: OutputReasoningItem
page_id: schema-outputreasoningitem-e39fd828
path: schemas
description: An output item containing reasoning
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputReasoningItem

An output item containing reasoning

```yaml
{"allOf": [{"$ref": "#/components/schemas/OutputItemReasoning"}, {"properties": {"content": {"items": {"$ref": "#/components/schemas/ReasoningTextContent"}, "type": ["array", "null"]}, "format": {"$ref": "#/components/schemas/ReasoningFormat"}, "signature": {"description": "A signature for the reasoning content, used for verification", "example": "EvcBCkgIChABGAIqQKkSDbRuVEQUk9qN1odC098l9SEj...", "type": ["string", "null"]}}, "type": "object"}], "description": "An output item containing reasoning", "example": {"content": [{"text": "First, we analyze the problem...", "type": "reasoning_text"}], "format": "anthropic-claude-v1", "id": "reasoning-123", "signature": "EvcBCkgIChABGAIqQKkSDbRuVEQUk9qN1odC098l9SEj...", "status": "completed", "summary": [{"text": "Analyzed the problem and found the optimal solution.", "type": "summary_text"}], "type": "reasoning"}}
```
