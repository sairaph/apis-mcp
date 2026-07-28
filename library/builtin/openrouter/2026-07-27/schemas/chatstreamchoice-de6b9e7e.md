---
title: ChatStreamChoice
page_id: schema-chatstreamchoice-de6b9e7e
path: schemas
description: Streaming completion choice chunk
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatStreamChoice

Streaming completion choice chunk

```yaml
{"description": "Streaming completion choice chunk", "example": {"delta": {"content": "Hello", "role": "assistant"}, "finish_reason": null, "index": 0}, "properties": {"delta": {"$ref": "#/components/schemas/ChatStreamDelta"}, "finish_reason": {"$ref": "#/components/schemas/ChatFinishReasonEnum"}, "index": {"description": "Choice index", "example": 0, "type": "integer"}, "logprobs": {"$ref": "#/components/schemas/ChatTokenLogprobs"}}, "required": ["delta", "finish_reason", "index"], "type": "object"}
```
