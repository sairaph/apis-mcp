---
title: CustomToolCallInputDoneEvent
page_id: schema-customtoolcallinputdoneevent-562ba600
path: schemas
description: Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CustomToolCallInputDoneEvent

Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseCustomToolCallInputDoneEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.", "example": {"input": "*** Begin Patch\n*** End Patch", "item_id": "item-1", "output_index": 0, "sequence_number": 6, "type": "response.custom_tool_call_input.done"}}
```
