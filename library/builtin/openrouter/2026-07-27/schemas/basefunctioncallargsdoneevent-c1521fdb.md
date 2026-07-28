---
title: BaseFunctionCallArgsDoneEvent
page_id: schema-basefunctioncallargsdoneevent-c1521fdb
path: schemas
description: Event emitted when function call arguments streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseFunctionCallArgsDoneEvent

Event emitted when function call arguments streaming is complete

```yaml
{"description": "Event emitted when function call arguments streaming is complete", "example": {"arguments": "{\"city\": \"San Francisco\", \"units\": \"celsius\"}", "item_id": "item-1", "name": "get_weather", "output_index": 0, "sequence_number": 6, "type": "response.function_call_arguments.done"}, "properties": {"arguments": {"type": "string"}, "item_id": {"type": "string"}, "name": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.function_call_arguments.done"], "type": "string"}}, "required": ["type", "item_id", "output_index", "name", "arguments", "sequence_number"], "type": "object"}
```
