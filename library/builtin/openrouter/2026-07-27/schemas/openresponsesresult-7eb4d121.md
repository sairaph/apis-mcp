---
title: OpenResponsesResult
page_id: schema-openresponsesresult-7eb4d121
path: schemas
description: Complete non-streaming response from the Responses API
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenResponsesResult

Complete non-streaming response from the Responses API

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseResponsesResult"}, {"properties": {"error_type": {"$ref": "#/components/schemas/ApiErrorType"}, "openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "output": {"items": {"$ref": "#/components/schemas/OutputItems"}, "type": "array"}, "service_tier": {"type": ["string", "null"]}, "text": {"$ref": "#/components/schemas/TextExtendedConfig"}, "usage": {"$ref": "#/components/schemas/Usage"}}, "type": "object"}], "description": "Complete non-streaming response from the Responses API", "example": {"completed_at": 1704067210, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [{"content": [{"annotations": [], "text": "Hello! How can I help you today?", "type": "output_text"}], "id": "msg-abc123", "role": "assistant", "status": "completed", "type": "message"}], "parallel_tool_calls": true, "presence_penalty": null, "status": "completed", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null, "usage": {"input_tokens": 10, "input_tokens_details": {"cached_tokens": 0}, "output_tokens": 25, "output_tokens_details": {"reasoning_tokens": 0}, "total_tokens": 35}}}
```
