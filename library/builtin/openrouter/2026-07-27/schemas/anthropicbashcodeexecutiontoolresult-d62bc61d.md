---
title: AnthropicBashCodeExecutionToolResult
page_id: schema-anthropicbashcodeexecutiontoolresult-d62bc61d
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicBashCodeExecutionToolResult

```yaml
{"example": {"content": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "bash_code_execution_result"}, "tool_use_id": "srvtoolu_01abc", "type": "bash_code_execution_tool_result"}, "properties": {"content": {"$ref": "#/components/schemas/AnthropicBashCodeExecutionContent"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["bash_code_execution_tool_result"], "type": "string"}}, "required": ["type", "content", "tool_use_id"], "type": "object"}
```
