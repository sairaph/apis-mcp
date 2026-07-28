---
title: MessagesDeltaEvent
page_id: schema-messagesdeltaevent-ed15e2d7
path: schemas
description: Event sent when the message metadata changes (e.g., stop_reason)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesDeltaEvent

Event sent when the message metadata changes (e.g., stop_reason)

```yaml
{"description": "Event sent when the message metadata changes (e.g., stop_reason)", "example": {"delta": {"container": null, "stop_details": null, "stop_reason": "end_turn", "stop_sequence": null}, "type": "message_delta", "usage": {"cache_creation_input_tokens": null, "cache_read_input_tokens": null, "input_tokens": null, "output_tokens": 15, "output_tokens_details": null, "server_tool_use": null}}, "properties": {"delta": {"properties": {"container": {"$ref": "#/components/schemas/AnthropicContainer"}, "stop_details": {"$ref": "#/components/schemas/AnthropicRefusalStopDetails"}, "stop_reason": {"$ref": "#/components/schemas/ORAnthropicStopReason"}, "stop_sequence": {"type": ["string", "null"]}}, "required": ["container", "stop_details", "stop_reason", "stop_sequence"], "type": "object"}, "type": {"enum": ["message_delta"], "type": "string"}, "usage": {"properties": {"cache_creation": {"$ref": "#/components/schemas/AnthropicCacheCreation"}, "cache_creation_input_tokens": {"type": ["integer", "null"]}, "cache_read_input_tokens": {"type": ["integer", "null"]}, "input_tokens": {"type": ["integer", "null"]}, "iterations": {"items": {"$ref": "#/components/schemas/AnthropicUsageIteration"}, "type": "array"}, "output_tokens": {"type": "integer"}, "output_tokens_details": {"$ref": "#/components/schemas/AnthropicOutputTokensDetails"}, "server_tool_use": {"properties": {"web_fetch_requests": {"type": "integer"}, "web_search_requests": {"type": "integer"}}, "required": ["web_search_requests", "web_fetch_requests"], "type": ["object", "null"]}}, "required": ["input_tokens", "output_tokens", "output_tokens_details", "cache_creation_input_tokens", "cache_read_input_tokens", "server_tool_use"], "type": "object"}}, "required": ["type", "delta", "usage"], "type": "object"}
```
