---
title: FailedEvent
page_id: schema-failedevent-9e216be2
path: schemas
description: Event emitted when a response has failed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FailedEvent

Event emitted when a response has failed

```yaml
{"description": "Event emitted when a response has failed", "example": {"response": {"completed_at": null, "created_at": 1704067200, "error": null, "frequency_penalty": null, "id": "resp-abc123", "incomplete_details": null, "instructions": null, "max_output_tokens": null, "metadata": null, "model": "gpt-4", "object": "response", "output": [], "parallel_tool_calls": true, "presence_penalty": null, "status": "failed", "temperature": null, "tool_choice": "auto", "tools": [], "top_p": null}, "sequence_number": 3, "type": "response.failed"}, "properties": {"response": {"$ref": "#/components/schemas/BaseResponsesResult"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.failed"], "type": "string"}}, "required": ["type", "response", "sequence_number"], "type": "object"}
```
