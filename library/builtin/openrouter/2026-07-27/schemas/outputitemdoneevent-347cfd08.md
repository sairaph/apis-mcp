---
title: OutputItemDoneEvent
page_id: schema-outputitemdoneevent-347cfd08
path: schemas
description: Event emitted when an output item is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemDoneEvent

Event emitted when an output item is complete

```yaml
{"description": "Event emitted when an output item is complete", "example": {"item": {"content": [{"annotations": [], "text": "Hello! How can I help you?", "type": "output_text"}], "id": "item-1", "role": "assistant", "status": "completed", "type": "message"}, "output_index": 0, "sequence_number": 8, "type": "response.output_item.done"}, "properties": {"item": {"discriminator": {"mapping": {"apply_patch_call": "#/components/schemas/OutputItemApplyPatchCall", "custom_tool_call": "#/components/schemas/OutputItemCustomToolCall", "file_search_call": "#/components/schemas/OutputItemFileSearchCall", "function_call": "#/components/schemas/OutputItemFunctionCall", "image_generation_call": "#/components/schemas/OutputItemImageGenerationCall", "message": "#/components/schemas/OutputMessage", "reasoning": "#/components/schemas/OutputItemReasoning", "web_search_call": "#/components/schemas/OutputItemWebSearchCall"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/OutputMessage"}, {"$ref": "#/components/schemas/OutputItemReasoning"}, {"$ref": "#/components/schemas/OutputItemFunctionCall"}, {"$ref": "#/components/schemas/OutputItemCustomToolCall"}, {"$ref": "#/components/schemas/OutputItemWebSearchCall"}, {"$ref": "#/components/schemas/OutputItemFileSearchCall"}, {"$ref": "#/components/schemas/OutputItemImageGenerationCall"}, {"$ref": "#/components/schemas/OutputItemApplyPatchCall"}]}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.output_item.done"], "type": "string"}}, "required": ["type", "output_index", "item", "sequence_number"], "type": "object"}
```
