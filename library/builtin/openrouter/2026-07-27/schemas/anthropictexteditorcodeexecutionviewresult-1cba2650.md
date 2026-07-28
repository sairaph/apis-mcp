---
title: AnthropicTextEditorCodeExecutionViewResult
page_id: schema-anthropictexteditorcodeexecutionviewresult-1cba2650
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextEditorCodeExecutionViewResult

```yaml
{"example": {"content": "file content", "file_type": "text", "num_lines": 10, "start_line": 1, "total_lines": 10, "type": "text_editor_code_execution_view_result"}, "properties": {"content": {"type": "string"}, "file_type": {"enum": ["text", "image", "pdf"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "num_lines": {"type": ["integer", "null"]}, "start_line": {"type": ["integer", "null"]}, "total_lines": {"type": ["integer", "null"]}, "type": {"enum": ["text_editor_code_execution_view_result"], "type": "string"}}, "required": ["content", "file_type", "num_lines", "start_line", "total_lines", "type"], "type": "object"}
```
