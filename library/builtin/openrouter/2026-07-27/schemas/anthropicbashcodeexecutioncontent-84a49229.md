---
title: AnthropicBashCodeExecutionContent
page_id: schema-anthropicbashcodeexecutioncontent-84a49229
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicBashCodeExecutionContent

```yaml
{"discriminator": {"mapping": {"bash_code_execution_result": "#/components/schemas/AnthropicBashCodeExecutionResult", "bash_code_execution_tool_result_error": "#/components/schemas/AnthropicBashCodeExecutionToolResultError"}, "propertyName": "type"}, "example": {"content": [], "return_code": 0, "stderr": "", "stdout": "Hello", "type": "bash_code_execution_result"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicBashCodeExecutionToolResultError"}, {"$ref": "#/components/schemas/AnthropicBashCodeExecutionResult"}]}
```
