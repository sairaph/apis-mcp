---
title: InProgressEvent
page_id: schema-inprogressevent-5d1f2706
path: schemas
description: Event emitted when a response is in progress
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InProgressEvent

Event emitted when a response is in progress

```yaml
{"description": "Event emitted when a response is in progress", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [], "parallel_tool_calls": true, "presence_penalty": null, "status": "in_progress", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 1, "type": "response.in_progress"}, "properties": {"response": {"$ref": "#/components/schemas/BaseResponsesResult"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.in_progress"], "type": "string"}}, "required": ["type", "response", "sequence_number"], "type": "object"}
```
