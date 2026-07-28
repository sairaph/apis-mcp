---
title: AnthropicCodeExecutionContent
page_id: schema-anthropiccodeexecutioncontent-1c996029
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCodeExecutionContent

```yaml
{"discriminator": {"mapping": {"code_execution_result": "#/components/schemas/AnthropicCodeExecutionResult", "code_execution_tool_result_error": "#/components/schemas/AnthropicCodeExecutionToolResultError", "encrypted_code_execution_result": "#/components/schemas/AnthropicEncryptedCodeExecutionResult"}, "propertyName": "type"}, "example": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "code_execution_result"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicCodeExecutionToolResultError"}, {"$ref": "#/components/schemas/AnthropicCodeExecutionResult"}, {"$ref": "#/components/schemas/AnthropicEncryptedCodeExecutionResult"}]}
```
