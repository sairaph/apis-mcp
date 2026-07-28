---
title: FunctionCallArgsDeltaEvent
page_id: schema-functioncallargsdeltaevent-36a9a66b
path: schemas
description: Event emitted when function call arguments are being streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FunctionCallArgsDeltaEvent

Event emitted when function call arguments are being streamed

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseFunctionCallArgsDeltaEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when function call arguments are being streamed", "example": {"delta": "{\"city\": \"...\"}", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.function_call_arguments.delta"}}
```
