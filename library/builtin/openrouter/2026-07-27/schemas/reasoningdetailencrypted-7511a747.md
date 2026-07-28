---
title: ReasoningDetailEncrypted
page_id: schema-reasoningdetailencrypted-7511a747
path: schemas
description: Reasoning detail encrypted schema
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDetailEncrypted

Reasoning detail encrypted schema

```yaml
{"description": "Reasoning detail encrypted schema", "example": {"data": "encrypted data", "type": "reasoning.encrypted"}, "properties": {"data": {"type": "string"}, "format": {"$ref": "#/components/schemas/ReasoningFormat"}, "id": {"type": ["string", "null"]}, "index": {"type": "integer"}, "type": {"enum": ["reasoning.encrypted"], "type": "string"}}, "required": ["type", "data"], "type": "object"}
```
