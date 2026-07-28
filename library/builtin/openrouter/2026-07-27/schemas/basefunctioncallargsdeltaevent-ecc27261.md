---
title: BaseFunctionCallArgsDeltaEvent
page_id: schema-basefunctioncallargsdeltaevent-ecc27261
path: schemas
description: Event emitted when function call arguments are being streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseFunctionCallArgsDeltaEvent

Event emitted when function call arguments are being streamed

```yaml
{"description": "Event emitted when function call arguments are being streamed", "example": {"delta": "{\"city\": \"...\"}", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.function_call_arguments.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.function_call_arguments.delta"], "type": "string"}}, "required": ["type", "item_id", "output_index", "delta", "sequence_number"], "type": "object"}
```
