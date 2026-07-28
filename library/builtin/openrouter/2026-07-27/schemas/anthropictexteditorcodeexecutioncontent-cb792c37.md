---
title: AnthropicTextEditorCodeExecutionContent
page_id: schema-anthropictexteditorcodeexecutioncontent-cb792c37
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextEditorCodeExecutionContent

```yaml
{"discriminator": {"mapping": {"text_editor_code_execution_create_result": "#/components/schemas/AnthropicTextEditorCodeExecutionCreateResult", "text_editor_code_execution_str_replace_result": "#/components/schemas/AnthropicTextEditorCodeExecutionStrReplaceResult", "text_editor_code_execution_tool_result_error": "#/components/schemas/AnthropicTextEditorCodeExecutionToolResultError", "text_editor_code_execution_view_result": "#/components/schemas/AnthropicTextEditorCodeExecutionViewResult"}, "propertyName": "type"}, "example": {"content": "file content", "file_type": "text", "num_lines": 10, "start_line": 1, "total_lines": 10, "type": "text_editor_code_execution_view_result"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionToolResultError"}, {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionViewResult"}, {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionCreateResult"}, {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionStrReplaceResult"}]}
```
