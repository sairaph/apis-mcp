---
title: ReasoningDoneEvent
page_id: schema-reasoningdoneevent-d4809d33
path: schemas
description: Event emitted when reasoning text streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDoneEvent

Event emitted when reasoning text streaming is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseReasoningDoneEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when reasoning text streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "sequence_number": 6, "text": "First, we need to identify the key components and then combine them logically.", "type": "response.reasoning_text.done"}}
```
