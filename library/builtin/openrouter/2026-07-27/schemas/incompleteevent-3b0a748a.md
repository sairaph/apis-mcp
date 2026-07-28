---
title: IncompleteEvent
page_id: schema-incompleteevent-3b0a748a
path: schemas
description: Event emitted when a response is incomplete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# IncompleteEvent

Event emitted when a response is incomplete

```yaml
{"description": "Event emitted when a response is incomplete", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [], "parallel_tool_calls": true, "presence_penalty": null, "status": "incomplete", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 5, "type": "response.incomplete"}, "properties": {"response": {"$ref": "#/components/schemas/BaseResponsesResult"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.incomplete"], "type": "string"}}, "required": ["type", "response", "sequence_number"], "type": "object"}
```
