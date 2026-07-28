---
title: CreatedEvent
page_id: schema-createdevent-da891fa1
path: schemas
description: Event emitted when a response is created
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CreatedEvent

Event emitted when a response is created

```yaml
{"description": "Event emitted when a response is created", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [], "parallel_tool_calls": true, "presence_penalty": null, "status": "in_progress", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 0, "type": "response.created"}, "properties": {"response": {"$ref": "#/components/schemas/BaseResponsesResult"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.created"], "type": "string"}}, "required": ["type", "response", "sequence_number"], "type": "object"}
```
