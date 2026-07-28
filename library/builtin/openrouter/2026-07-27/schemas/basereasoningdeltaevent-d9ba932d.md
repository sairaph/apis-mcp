---
title: BaseReasoningDeltaEvent
page_id: schema-basereasoningdeltaevent-d9ba932d
path: schemas
description: Event emitted when reasoning text delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningDeltaEvent

Event emitted when reasoning text delta is streamed

```yaml
{"description": "Event emitted when reasoning text delta is streamed", "example": {"content_index": 0, "delta": "First, we need", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.reasoning_text.delta"}, "properties": {"content_index": {"type": "integer"}, "delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.reasoning_text.delta"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "delta", "sequence_number"], "type": "object"}
```
