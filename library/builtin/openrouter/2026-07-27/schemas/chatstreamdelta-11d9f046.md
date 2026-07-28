---
title: ChatStreamDelta
page_id: schema-chatstreamdelta-11d9f046
path: schemas
description: Delta changes in streaming response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatStreamDelta

Delta changes in streaming response

```yaml
{"description": "Delta changes in streaming response", "example": {"content": "Hello", "role": "assistant"}, "properties": {"audio": {"allOf": [{"$ref": "#/components/schemas/ChatAudioOutput"}, {"description": "Audio output data"}]}, "content": {"description": "Message content delta", "example": "Hello", "type": ["string", "null"]}, "reasoning": {"description": "Reasoning content delta", "example": "I need to", "type": ["string", "null"]}, "reasoning_details": {"$ref": "#/components/schemas/ChatStreamReasoningDetails"}, "refusal": {"description": "Refusal message delta", "example": null, "type": ["string", "null"]}, "role": {"description": "The role of the message author", "enum": ["assistant"], "example": "assistant", "type": "string"}, "tool_calls": {"description": "Tool calls delta", "items": {"$ref": "#/components/schemas/ChatStreamToolCall"}, "type": "array"}}, "type": "object"}
```
