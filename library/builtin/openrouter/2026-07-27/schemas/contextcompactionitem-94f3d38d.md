---
title: ContextCompactionItem
page_id: schema-contextcompactionitem-94f3d38d
path: schemas
description: A context compaction marker with an optional encrypted summary
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContextCompactionItem

A context compaction marker with an optional encrypted summary

```yaml
{"additionalProperties": {}, "description": "A context compaction marker with an optional encrypted summary", "example": {"encrypted_content": "enc_abc123...", "type": "context_compaction"}, "properties": {"encrypted_content": {"type": ["string", "null"]}, "id": {"type": ["string", "null"]}, "type": {"enum": ["context_compaction"], "type": "string"}}, "required": ["type"], "type": "object"}
```
