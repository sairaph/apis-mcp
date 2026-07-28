---
title: BaseMessagesResult
page_id: schema-basemessagesresult-66b71d56
path: schemas
description: Base Anthropic Messages API response before OpenRouter extensions
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseMessagesResult

Base Anthropic Messages API response before OpenRouter extensions

```yaml
{"description": "Base Anthropic Messages API response before OpenRouter extensions", "example": {"container": null, "content": [{"citations": [], "text": "Hello!", "type": "text"}], "id": "msg_01XFDUDYJgAACzvnptvVoYEL", "model": "claude-sonnet-4-5-20250929", "role": "assistant", "stop_details": null, "stop_reason": "end_turn", "stop_sequence": null, "type": "message", "usage": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 12, "output_tokens": 8, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}}, "properties": {"container": {"$ref": "#/components/schemas/AnthropicContainer"}, "content": {"items": {"$ref": "#/components/schemas/ORAnthropicContentBlock"}, "type": "array"}, "id": {"type": "string"}, "model": {"type": "string"}, "role": {"enum": ["assistant"], "type": "string"}, "stop_details": {"$ref": "#/components/schemas/AnthropicRefusalStopDetails"}, "stop_reason": {"$ref": "#/components/schemas/ORAnthropicStopReason"}, "stop_sequence": {"type": ["string", "null"]}, "type": {"enum": ["message"], "type": "string"}, "usage": {"allOf": [{"$ref": "#/components/schemas/AnthropicUsage"}, {"properties": {"iterations": {"items": {"$ref": "#/components/schemas/AnthropicUsageIteration"}, "type": "array"}, "speed": {"$ref": "#/components/schemas/AnthropicSpeed"}}, "type": "object"}], "example": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 100, "output_tokens": 50, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}}}, "required": ["id", "type", "role", "container", "content", "model", "stop_reason", "stop_details", "stop_sequence", "usage"], "type": "object"}
```
