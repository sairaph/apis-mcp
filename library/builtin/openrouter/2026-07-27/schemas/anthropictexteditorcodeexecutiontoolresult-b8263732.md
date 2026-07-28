---
title: AnthropicTextEditorCodeExecutionToolResult
page_id: schema-anthropictexteditorcodeexecutiontoolresult-b8263732
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextEditorCodeExecutionToolResult

```yaml
{"example": {"content": {"content": "file content", "file_type": "text", "num_lines": 10, "start_line": 1, "total_lines": 10, "type": "text_editor_code_execution_view_result"}, "tool_use_id": "srvtoolu_01abc", "type": "text_editor_code_execution_tool_result"}, "properties": {"content": {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionContent"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["text_editor_code_execution_tool_result"], "type": "string"}}, "required": ["type", "content", "tool_use_id"], "type": "object"}
```
