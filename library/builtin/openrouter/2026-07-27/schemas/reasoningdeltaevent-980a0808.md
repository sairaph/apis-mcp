---
title: ReasoningDeltaEvent
page_id: schema-reasoningdeltaevent-980a0808
path: schemas
description: Event emitted when reasoning text delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDeltaEvent

Event emitted when reasoning text delta is streamed

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseReasoningDeltaEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when reasoning text delta is streamed", "example": {"content_index": 0, "delta": "First, we need", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.reasoning_text.delta"}}
```
