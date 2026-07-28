---
title: AnthropicEncryptedCodeExecutionResult
page_id: schema-anthropicencryptedcodeexecutionresult-5a1f2471
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicEncryptedCodeExecutionResult

```yaml
{"example": {"content": [], "encrypted_stdout": "enc_stdout", "return_code": 0, "stderr": "", "type": "encrypted_code_execution_result"}, "properties": {"content": {"items": {"$ref": "#/components/schemas/AnthropicCodeExecutionOutput"}, "type": "array"}, "encrypted_stdout": {"type": "string"}, "return_code": {"type": "integer"}, "stderr": {"type": "string"}, "type": {"enum": ["encrypted_code_execution_result"], "type": "string"}}, "required": ["content", "encrypted_stdout", "return_code", "stderr", "type"], "type": "object"}
```
