---
title: MessagesStopEvent
page_id: schema-messagesstopevent-a9a77d7c
path: schemas
description: Event sent when the message is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesStopEvent

Event sent when the message is complete

```yaml
{"description": "Event sent when the message is complete", "example": {"type": "message_stop"}, "properties": {"openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "type": {"enum": ["message_stop"], "type": "string"}}, "required": ["type"], "type": "object"}
```
