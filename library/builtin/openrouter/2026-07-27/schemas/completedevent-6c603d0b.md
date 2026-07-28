---
title: CompletedEvent
page_id: schema-completedevent-6c603d0b
path: schemas
description: Event emitted when a response has completed successfully
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CompletedEvent

Event emitted when a response has completed successfully

```yaml
{"description": "Event emitted when a response has completed successfully", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [{"content": [{"annotations": [], "text": "Hello! How can I help you?", "type": "output_text"}], "id": "item-1", "role": "assistant", "status": "completed", "type": "message"}], "parallel_tool_calls": true, "presence_penalty": null, "status": "completed", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 10, "type": "response.completed"}, "properties": {"response": {"$ref": "#/components/schemas/BaseResponsesResult"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.completed"], "type": "string"}}, "required": ["type", "response", "sequence_number"], "type": "object"}
```
