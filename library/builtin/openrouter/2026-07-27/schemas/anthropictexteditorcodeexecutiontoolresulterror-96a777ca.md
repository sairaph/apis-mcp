---
title: AnthropicTextEditorCodeExecutionToolResultError
page_id: schema-anthropictexteditorcodeexecutiontoolresulterror-96a777ca
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextEditorCodeExecutionToolResultError

```yaml
{"example": {"error_code": "unavailable", "error_message": null, "type": "text_editor_code_execution_tool_result_error"}, "properties": {"error_code": {"enum": ["invalid_tool_input", "unavailable", "too_many_requests", "execution_time_exceeded", "file_not_found"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "error_message": {"type": ["string", "null"]}, "type": {"enum": ["text_editor_code_execution_tool_result_error"], "type": "string"}}, "required": ["error_code", "error_message", "type"], "type": "object"}
```
