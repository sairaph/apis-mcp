---
title: OpenAIResponseFunctionToolCall
page_id: schema-openairesponsefunctiontoolcall-5673bf17
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponseFunctionToolCall

```yaml
{"example": {"arguments": "{\"location\":\"San Francisco\"}", "call_id": "call-abc123", "id": "fc-abc123", "name": "get_weather", "status": "completed", "type": "function_call"}, "properties": {"arguments": {"type": "string"}, "call_id": {"type": "string"}, "id": {"type": "string"}, "name": {"type": "string"}, "namespace": {"description": "Namespace qualifier for tools registered as part of a namespace tool group (e.g. an MCP server)", "type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["function_call"], "type": "string"}}, "required": ["type", "call_id", "name", "arguments"], "type": "object"}
```
