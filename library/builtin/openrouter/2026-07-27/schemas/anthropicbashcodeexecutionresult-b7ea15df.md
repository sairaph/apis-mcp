---
title: AnthropicBashCodeExecutionResult
page_id: schema-anthropicbashcodeexecutionresult-b7ea15df
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicBashCodeExecutionResult

```yaml
{"example": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "bash_code_execution_result"}, "properties": {"content": {"items": {"$ref": "#/components/schemas/AnthropicBashCodeExecutionOutput"}, "type": "array"}, "return_code": {"type": "integer"}, "stderr": {"type": "string"}, "stdout": {"type": "string"}, "type": {"enum": ["bash_code_execution_result"], "type": "string"}}, "required": ["content", "return_code", "stderr", "stdout", "type"], "type": "object"}
```
