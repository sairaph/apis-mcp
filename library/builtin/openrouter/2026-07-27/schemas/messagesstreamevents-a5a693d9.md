---
title: MessagesStreamEvents
page_id: schema-messagesstreamevents-a5a693d9
path: schemas
description: Union of all possible streaming events
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesStreamEvents

Union of all possible streaming events

```yaml
{"description": "Union of all possible streaming events", "discriminator": {"mapping": {"content_block_delta": "#/components/schemas/MessagesContentBlockDeltaEvent", "content_block_start": "#/components/schemas/MessagesContentBlockStartEvent", "content_block_stop": "#/components/schemas/MessagesContentBlockStopEvent", "error": "#/components/schemas/MessagesErrorEvent", "message_delta": "#/components/schemas/MessagesDeltaEvent", "message_start": "#/components/schemas/MessagesStartEvent", "message_stop": "#/components/schemas/MessagesStopEvent", "ping": "#/components/schemas/MessagesPingEvent"}, "propertyName": "type"}, "example": {"delta": {"text": "Hello", "type": "text_delta"}, "index": 0, "type": "content_block_delta"}, "oneOf": [{"$ref": "#/components/schemas/MessagesStartEvent"}, {"$ref": "#/components/schemas/MessagesDeltaEvent"}, {"$ref": "#/components/schemas/MessagesStopEvent"}, {"$ref": "#/components/schemas/MessagesContentBlockStartEvent"}, {"$ref": "#/components/schemas/MessagesContentBlockDeltaEvent"}, {"$ref": "#/components/schemas/MessagesContentBlockStopEvent"}, {"$ref": "#/components/schemas/MessagesPingEvent"}, {"$ref": "#/components/schemas/MessagesErrorEvent"}]}
```
