---
title: OutputItemReasoning
page_id: schema-outputitemreasoning-d65565dd
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemReasoning

```yaml
{"example": {"id": "reasoning-abc123", "summary": [{"text": "Analyzed the problem using first principles", "type": "summary_text"}], "type": "reasoning"}, "properties": {"content": {"items": {"$ref": "#/components/schemas/ReasoningTextContent"}, "type": "array"}, "encrypted_content": {"type": ["string", "null"]}, "id": {"type": "string"}, "status": {"anyOf": [{"enum": ["completed"], "type": "string"}, {"enum": ["incomplete"], "type": "string"}, {"enum": ["in_progress"], "type": "string"}]}, "summary": {"items": {"$ref": "#/components/schemas/ReasoningSummaryText"}, "type": "array"}, "type": {"enum": ["reasoning"], "type": "string"}}, "required": ["type", "id", "summary"], "type": "object"}
```
