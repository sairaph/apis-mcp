---
title: AnthropicCodeExecutionToolResult
page_id: schema-anthropiccodeexecutiontoolresult-12d2e201
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCodeExecutionToolResult

```yaml
{"example": {"content": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "code_execution_result"}, "tool_use_id": "srvtoolu_01abc", "type": "code_execution_tool_result"}, "properties": {"content": {"$ref": "#/components/schemas/AnthropicCodeExecutionContent"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["code_execution_tool_result"], "type": "string"}}, "required": ["type", "content", "tool_use_id"], "type": "object"}
```
