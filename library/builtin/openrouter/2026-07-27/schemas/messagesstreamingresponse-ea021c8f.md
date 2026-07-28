---
title: MessagesStreamingResponse
page_id: schema-messagesstreamingresponse-ea021c8f
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesStreamingResponse

```yaml
{"example": {"data": {"delta": {"text": "Hello", "type": "text_delta"}, "index": 0, "type": "content_block_delta"}, "event": "content_block_delta"}, "properties": {"data": {"$ref": "#/components/schemas/MessagesStreamEvents"}, "event": {"type": "string"}}, "required": ["event", "data"], "type": "object"}
```
