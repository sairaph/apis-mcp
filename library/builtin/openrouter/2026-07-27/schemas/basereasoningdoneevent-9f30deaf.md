---
title: BaseReasoningDoneEvent
page_id: schema-basereasoningdoneevent-9f30deaf
path: schemas
description: Event emitted when reasoning text streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningDoneEvent

Event emitted when reasoning text streaming is complete

```yaml
{"description": "Event emitted when reasoning text streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "sequence_number": 6, "text": "First, we need to identify the key components and then combine them logically.", "type": "response.reasoning_text.done"}, "properties": {"content_index": {"type": "integer"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "text": {"type": "string"}, "type": {"enum": ["response.reasoning_text.done"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "text", "sequence_number"], "type": "object"}
```
