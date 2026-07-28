---
title: ReasoningDetailServerToolCall
page_id: schema-reasoningdetailservertoolcall-d800cc7b
path: schemas
description: Record of an OpenRouter server-tool invocation (e.g. openrouter:fusion), carried in reasoning_details so a prior tool call can be rehydrated into a later turn of the same conversation.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDetailServerToolCall

Record of an OpenRouter server-tool invocation (e.g. openrouter:fusion), carried in reasoning_details so a prior tool call can be rehydrated into a later turn of the same conversation.

```yaml
{"description": "Record of an OpenRouter server-tool invocation (e.g. openrouter:fusion), carried in reasoning_details so a prior tool call can be rehydrated into a later turn of the same conversation.", "example": {"arguments": "{\"prompt\":\"Compare carbon tax proposals\"}", "result": "{\"status\":\"ok\",\"models\":[\"openai/gpt-4o\"]}", "tool_call_id": "call_abc123", "tool_name": "openrouter:fusion", "type": "reasoning.server_tool_call"}, "properties": {"arguments": {"type": "string"}, "format": {"$ref": "#/components/schemas/ReasoningFormat"}, "id": {"type": ["string", "null"]}, "index": {"type": "integer"}, "result": {"type": "string"}, "tool_call_id": {"type": ["string", "null"]}, "tool_name": {"type": "string"}, "type": {"enum": ["reasoning.server_tool_call"], "type": "string"}}, "required": ["type", "tool_name", "arguments", "result"], "type": "object"}
```
