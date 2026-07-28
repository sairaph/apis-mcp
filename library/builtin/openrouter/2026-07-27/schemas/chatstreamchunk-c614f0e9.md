---
title: ChatStreamChunk
page_id: schema-chatstreamchunk-c614f0e9
path: schemas
description: Streaming chat completion chunk
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatStreamChunk

Streaming chat completion chunk

```yaml
{"description": "Streaming chat completion chunk", "example": {"choices": [{"delta": {"content": "Hello", "role": "assistant"}, "finish_reason": null, "index": 0}], "created": 1677652288, "id": "chatcmpl-123", "model": "openai/gpt-4", "object": "chat.completion.chunk"}, "properties": {"choices": {"description": "List of streaming chunk choices", "items": {"$ref": "#/components/schemas/ChatStreamChoice"}, "type": "array"}, "created": {"description": "Unix timestamp of creation", "example": 1677652288, "type": "integer"}, "error": {"description": "Error information", "example": {"code": 429, "message": "Rate limit exceeded", "metadata": {"error_type": "rate_limit_exceeded"}}, "properties": {"code": {"description": "Error code", "example": 429, "format": "int32", "type": "integer"}, "message": {"description": "Error message", "example": "Rate limit exceeded", "type": "string"}, "metadata": {"description": "Structured error metadata", "properties": {"error_type": {"$ref": "#/components/schemas/ApiErrorType"}, "provider_code": {"description": "Upstream provider-specific error code, when available", "type": "string"}}, "required": ["error_type"], "type": "object"}}, "required": ["message", "code"], "type": "object"}, "id": {"description": "Unique chunk identifier", "example": "chatcmpl-123", "type": "string"}, "model": {"description": "Model used for completion", "example": "openai/gpt-4", "type": "string"}, "object": {"enum": ["chat.completion.chunk"], "type": "string"}, "openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "service_tier": {"description": "The service tier used by the upstream provider for this request", "example": "default", "type": ["string", "null"]}, "system_fingerprint": {"description": "System fingerprint", "example": "fp_44709d6fcb", "type": "string"}, "usage": {"$ref": "#/components/schemas/ChatUsage"}}, "required": ["id", "choices", "created", "model", "object"], "type": "object", "x-speakeasy-entity": "ChatStreamChunk"}
```
