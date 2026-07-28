---
title: ChatAssistantMessage
page_id: schema-chatassistantmessage-310ee8f1
path: schemas
description: Assistant message for requests and responses
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatAssistantMessage

Assistant message for requests and responses

```yaml
{"description": "Assistant message for requests and responses", "example": {"content": "The capital of France is Paris.", "model": "openai/gpt-4o", "role": "assistant"}, "properties": {"audio": {"$ref": "#/components/schemas/ChatAudioOutput"}, "content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/ChatContentItems"}, "type": "array"}, {"type": "null"}], "description": "Assistant message content"}, "images": {"$ref": "#/components/schemas/ChatAssistantImages"}, "model": {"description": "Model that generated this assistant message", "example": "openai/gpt-4o", "type": "string"}, "name": {"description": "Optional name for the assistant", "type": "string"}, "reasoning": {"description": "Reasoning output", "type": ["string", "null"]}, "reasoning_details": {"$ref": "#/components/schemas/ChatReasoningDetails"}, "refusal": {"description": "Refusal message if content was refused", "type": ["string", "null"]}, "role": {"enum": ["assistant"], "type": "string"}, "tool_calls": {"description": "Tool calls made by the assistant", "items": {"$ref": "#/components/schemas/ChatToolCall"}, "type": "array"}}, "required": ["role"], "type": "object"}
```
