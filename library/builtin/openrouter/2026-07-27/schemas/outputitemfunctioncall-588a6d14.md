---
title: OutputItemFunctionCall
page_id: schema-outputitemfunctioncall-588a6d14
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemFunctionCall

```yaml
{"example": {"arguments": "{\"location\":\"San Francisco\",\"unit\":\"celsius\"}", "call_id": "call-abc123", "id": "call-abc123", "name": "get_weather", "type": "function_call"}, "properties": {"arguments": {"type": "string"}, "call_id": {"type": "string"}, "id": {"type": "string"}, "name": {"type": "string"}, "namespace": {"description": "Namespace qualifier for tools registered as part of a namespace tool group (e.g. an MCP server)", "type": "string"}, "status": {"anyOf": [{"enum": ["completed"], "type": "string"}, {"enum": ["incomplete"], "type": "string"}, {"enum": ["in_progress"], "type": "string"}]}, "type": {"enum": ["function_call"], "type": "string"}}, "required": ["type", "name", "arguments", "call_id"], "type": "object"}
```
