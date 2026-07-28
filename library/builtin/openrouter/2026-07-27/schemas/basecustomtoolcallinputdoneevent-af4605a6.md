---
title: BaseCustomToolCallInputDoneEvent
page_id: schema-basecustomtoolcallinputdoneevent-af4605a6
path: schemas
description: Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseCustomToolCallInputDoneEvent

Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.

```yaml
{"description": "Event emitted when a custom tool call's freeform input streaming is complete. Mirrors `response.function_call_arguments.done` but for `custom` tools.", "example": {"input": "*** Begin Patch\n*** End Patch", "item_id": "item-1", "output_index": 0, "sequence_number": 6, "type": "response.custom_tool_call_input.done"}, "properties": {"input": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.custom_tool_call_input.done"], "type": "string"}}, "required": ["type", "item_id", "output_index", "input", "sequence_number"], "type": "object"}
```
