---
title: CustomToolCallOutputItem
page_id: schema-customtoolcalloutputitem-64865460
path: schemas
description: The output from a custom (freeform-grammar) tool call execution. Mirrors `function_call_output` but is matched to a `custom_tool_call` rather than a `function_call`.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CustomToolCallOutputItem

The output from a custom (freeform-grammar) tool call execution. Mirrors `function_call_output` but is matched to a `custom_tool_call` rather than a `function_call`.

```yaml
{"allOf": [{"$ref": "#/components/schemas/OpenAIResponseCustomToolCallOutput"}, {"properties": {"output": {"anyOf": [{"type": "string"}, {"items": {"oneOf": [{"$ref": "#/components/schemas/InputText"}, {"allOf": [{"$ref": "#/components/schemas/InputImage"}, {"properties": {}, "type": "object"}], "description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}}, {"$ref": "#/components/schemas/InputFile"}]}, "type": "array"}]}}, "type": "object"}], "description": "The output from a custom (freeform-grammar) tool call execution. Mirrors `function_call_output` but is matched to a `custom_tool_call` rather than a `function_call`.", "example": {"call_id": "call-abc123", "id": "ctco-abc123", "output": "patch applied successfully", "type": "custom_tool_call_output"}}
```
