---
title: CompactionItem
page_id: schema-compactionitem-46933c6d
path: schemas
description: A context compaction marker with encrypted summary
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CompactionItem

A context compaction marker with encrypted summary

```yaml
{"additionalProperties": {}, "description": "A context compaction marker with encrypted summary", "example": {"encrypted_content": "enc_abc123...", "type": "compaction"}, "properties": {"encrypted_content": {"type": "string"}, "id": {"type": ["string", "null"]}, "type": {"enum": ["compaction"], "type": "string"}}, "required": ["type", "encrypted_content"], "type": "object"}
```
