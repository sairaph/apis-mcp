---
title: BaseErrorEvent
page_id: schema-baseerrorevent-2e0cfe5d
path: schemas
description: Event emitted when an error occurs during streaming
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseErrorEvent

Event emitted when an error occurs during streaming

```yaml
{"description": "Event emitted when an error occurs during streaming", "example": {"code": "rate_limit_exceeded", "message": "Rate limit exceeded. Please try again later.", "param": null, "sequence_number": 2, "type": "error"}, "properties": {"code": {"type": ["string", "null"]}, "message": {"type": "string"}, "param": {"type": ["string", "null"]}, "sequence_number": {"type": "integer"}, "type": {"enum": ["error"], "type": "string"}}, "required": ["type", "code", "message", "param", "sequence_number"], "type": "object"}
```
