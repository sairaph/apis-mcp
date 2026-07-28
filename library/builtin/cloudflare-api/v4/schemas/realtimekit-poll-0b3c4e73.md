---
title: realtimekit_Poll
page_id: schema-realtimekit-poll-0b3c4e73
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_Poll

```yaml
{"type": "object", "properties": {"anonymous": {"type": "boolean"}, "created_by": {"type": "string"}, "hide_votes": {"type": "boolean"}, "id": {"description": "ID of the poll", "type": "string", "readOnly": true}, "options": {"description": "Answer options", "type": "array", "items": {"properties": {"count": {"type": "number"}, "text": {"description": "Text of the answer option", "type": "string"}, "votes": {"type": "array", "items": {"properties": {"id": {"type": "string"}, "name": {"type": "string"}}, "required": ["id", "name"], "type": "object"}}}, "required": ["text", "count", "votes"], "type": "object"}}, "question": {"description": "Question asked by the poll", "type": "string"}, "voted": {"type": "array", "items": {"type": "string"}}}, "required": ["id", "question", "options"]}
```
