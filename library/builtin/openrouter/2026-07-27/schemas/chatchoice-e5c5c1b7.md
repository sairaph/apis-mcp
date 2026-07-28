---
title: ChatChoice
page_id: schema-chatchoice-e5c5c1b7
path: schemas
description: Chat completion choice
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatChoice

Chat completion choice

```yaml
{"description": "Chat completion choice", "example": {"finish_reason": "stop", "index": 0, "logprobs": null, "message": {"content": "The capital of France is Paris.", "role": "assistant"}}, "properties": {"finish_reason": {"$ref": "#/components/schemas/ChatFinishReasonEnum"}, "index": {"description": "Choice index", "example": 0, "type": "integer"}, "logprobs": {"$ref": "#/components/schemas/ChatTokenLogprobs"}, "message": {"$ref": "#/components/schemas/ChatAssistantMessage"}}, "required": ["finish_reason", "index", "message"], "type": "object"}
```
