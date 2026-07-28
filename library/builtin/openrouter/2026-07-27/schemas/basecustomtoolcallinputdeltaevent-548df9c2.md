---
title: BaseCustomToolCallInputDeltaEvent
page_id: schema-basecustomtoolcallinputdeltaevent-548df9c2
path: schemas
description: Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseCustomToolCallInputDeltaEvent

Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.

```yaml
{"description": "Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.", "example": {"delta": "*** Begin Patch", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.custom_tool_call_input.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.custom_tool_call_input.delta"], "type": "string"}}, "required": ["type", "item_id", "output_index", "delta", "sequence_number"], "type": "object"}
```
