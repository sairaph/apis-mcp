---
title: ChatResult
page_id: schema-chatresult-8795d0a6
path: schemas
description: Chat completion response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatResult

Chat completion response

```yaml
{"description": "Chat completion response", "example": {"choices": [{"finish_reason": "stop", "index": 0, "message": {"content": "The capital of France is Paris.", "role": "assistant"}}], "created": 1677652288, "id": "chatcmpl-123", "model": "openai/gpt-4", "object": "chat.completion", "system_fingerprint": "fp_44709d6fcb", "usage": {"completion_tokens": 15, "prompt_tokens": 10, "total_tokens": 25}}, "properties": {"choices": {"description": "List of completion choices", "items": {"$ref": "#/components/schemas/ChatChoice"}, "type": "array"}, "created": {"description": "Unix timestamp of creation", "example": 1677652288, "type": "integer"}, "id": {"description": "Unique completion identifier", "example": "chatcmpl-123", "type": "string"}, "model": {"description": "Model used for completion", "example": "openai/gpt-4", "type": "string"}, "object": {"enum": ["chat.completion"], "type": "string"}, "openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "service_tier": {"description": "The service tier used by the upstream provider for this request", "example": "default", "type": ["string", "null"]}, "system_fingerprint": {"description": "System fingerprint", "example": "fp_44709d6fcb", "type": ["string", "null"]}, "usage": {"$ref": "#/components/schemas/ChatUsage"}}, "required": ["id", "choices", "created", "model", "object", "system_fingerprint"], "type": "object"}
```
