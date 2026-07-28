---
title: ErrorEvent
page_id: schema-errorevent-70362dd7
path: schemas
description: Event emitted when an error occurs during streaming
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ErrorEvent

Event emitted when an error occurs during streaming

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseErrorEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when an error occurs during streaming", "example": {"code": "rate_limit_exceeded", "message": "Rate limit exceeded. Please try again later.", "param": null, "sequence_number": 2, "type": "error"}}
```
