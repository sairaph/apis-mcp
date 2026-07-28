---
title: OutputItemAddedEvent
page_id: schema-outputitemaddedevent-1e1b3505
path: schemas
description: Event emitted when a new output item is added to the response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemAddedEvent

Event emitted when a new output item is added to the response

```yaml
{"description": "Event emitted when a new output item is added to the response", "example": {"item": {"content": [], "id": "item-1", "role": "assistant", "status": "in_progress", "type": "message"}, "output_index": 0, "sequence_number": 2, "type": "response.output_item.added"}, "properties": {"item": {"discriminator": {"mapping": {"apply_patch_call": "#/components/schemas/OutputItemApplyPatchCall", "custom_tool_call": "#/components/schemas/OutputItemCustomToolCall", "file_search_call": "#/components/schemas/OutputItemFileSearchCall", "function_call": "#/components/schemas/OutputItemFunctionCall", "image_generation_call": "#/components/schemas/OutputItemImageGenerationCall", "message": "#/components/schemas/OutputMessage", "reasoning": "#/components/schemas/OutputItemReasoning", "web_search_call": "#/components/schemas/OutputItemWebSearchCall"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/OutputMessage"}, {"$ref": "#/components/schemas/OutputItemReasoning"}, {"$ref": "#/components/schemas/OutputItemFunctionCall"}, {"$ref": "#/components/schemas/OutputItemCustomToolCall"}, {"$ref": "#/components/schemas/OutputItemWebSearchCall"}, {"$ref": "#/components/schemas/OutputItemFileSearchCall"}, {"$ref": "#/components/schemas/OutputItemImageGenerationCall"}, {"$ref": "#/components/schemas/OutputItemApplyPatchCall"}]}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.output_item.added"], "type": "string"}}, "required": ["type", "output_index", "item", "sequence_number"], "type": "object"}
```
