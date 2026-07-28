---
title: MessagesResult
page_id: schema-messagesresult-3e0ad4e5
path: schemas
description: Non-streaming response from the Anthropic Messages API with OpenRouter extensions
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesResult

Non-streaming response from the Anthropic Messages API with OpenRouter extensions

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseMessagesResult"}, {"properties": {"context_management": {"properties": {"applied_edits": {"items": {"additionalProperties": {}, "properties": {"type": {"type": "string"}}, "required": ["type"], "type": "object"}, "type": "array"}}, "required": ["applied_edits"], "type": ["object", "null"]}, "openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "provider": {"$ref": "#/components/schemas/ProviderName"}, "usage": {"allOf": [{"$ref": "#/components/schemas/AnthropicUsage"}, {"properties": {"cost": {"format": "double", "type": ["number", "null"]}, "cost_details": {"$ref": "#/components/schemas/CostDetails"}, "is_byok": {"type": "boolean"}, "iterations": {"items": {"$ref": "#/components/schemas/AnthropicUsageIteration"}, "type": "array"}, "service_tier": {"type": ["string", "null"]}, "speed": {"$ref": "#/components/schemas/AnthropicSpeed"}}, "type": "object"}], "example": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 100, "output_tokens": 50, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}}}, "type": "object"}], "description": "Non-streaming response from the Anthropic Messages API with OpenRouter extensions", "example": {"container": null, "content": [{"citations": [], "text": "Hello! I'm doing well, thank you for asking.", "type": "text"}], "id": "msg_01XFDUDYJgAACzvnptvVoYEL", "model": "claude-sonnet-4-5-20250929", "role": "assistant", "stop_details": null, "stop_reason": "end_turn", "stop_sequence": null, "type": "message", "usage": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 12, "output_tokens": 15, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}}}
```
