---
title: AnthropicCodeExecutionResult
page_id: schema-anthropiccodeexecutionresult-d3ad2a43
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCodeExecutionResult

```yaml
{"example": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "code_execution_result"}, "properties": {"content": {"items": {"$ref": "#/components/schemas/AnthropicCodeExecutionOutput"}, "type": "array"}, "return_code": {"type": "integer"}, "stderr": {"type": "string"}, "stdout": {"type": "string"}, "type": {"enum": ["code_execution_result"], "type": "string"}}, "required": ["content", "return_code", "stderr", "stdout", "type"], "type": "object"}
```
