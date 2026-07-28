---
title: FunctionCallArgsDoneEvent
page_id: schema-functioncallargsdoneevent-08cda19b
path: schemas
description: Event emitted when function call arguments streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FunctionCallArgsDoneEvent

Event emitted when function call arguments streaming is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseFunctionCallArgsDoneEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when function call arguments streaming is complete", "example": {"arguments": "{\"city\": \"San Francisco\", \"units\": \"celsius\"}", "item_id": "item-1", "name": "get_weather", "output_index": 0, "sequence_number": 6, "type": "response.function_call_arguments.done"}}
```
