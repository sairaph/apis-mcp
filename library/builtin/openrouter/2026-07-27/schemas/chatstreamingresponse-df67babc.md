---
title: ChatStreamingResponse
page_id: schema-chatstreamingresponse-df67babc
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatStreamingResponse

```yaml
{"example": {"data": {"choices": [{"delta": {"content": "Hello", "role": "assistant"}, "finish_reason": null, "index": 0}], "created": 1677652288, "id": "chatcmpl-123", "model": "openai/gpt-4", "object": "chat.completion.chunk"}}, "properties": {"data": {"$ref": "#/components/schemas/ChatStreamChunk"}}, "required": ["data"], "type": "object"}
```
