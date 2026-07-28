---
title: CustomToolCallInputDeltaEvent
page_id: schema-customtoolcallinputdeltaevent-4f29cb0b
path: schemas
description: Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CustomToolCallInputDeltaEvent

Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseCustomToolCallInputDeltaEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a custom tool call's freeform input is being streamed. Mirrors `response.function_call_arguments.delta` but for `custom` tools whose input is opaque text rather than JSON arguments.", "example": {"delta": "*** Begin Patch", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.custom_tool_call_input.delta"}}
```
