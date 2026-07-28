---
title: StreamEventsResponseCompleted
page_id: schema-streameventsresponsecompleted-6b5a1dcc
path: schemas
description: Event emitted when a response has completed successfully
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StreamEventsResponseCompleted

Event emitted when a response has completed successfully

```yaml
{"allOf": [{"$ref": "#/components/schemas/CompletedEvent"}, {"properties": {"response": {"$ref": "#/components/schemas/OpenResponsesResult"}}, "type": "object"}], "description": "Event emitted when a response has completed successfully", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [], "parallel_tool_calls": true, "presence_penalty": null, "status": "completed", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 10, "type": "response.completed"}}
```
