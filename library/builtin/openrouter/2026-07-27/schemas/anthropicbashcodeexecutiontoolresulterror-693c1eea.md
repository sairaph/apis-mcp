---
title: AnthropicBashCodeExecutionToolResultError
page_id: schema-anthropicbashcodeexecutiontoolresulterror-693c1eea
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicBashCodeExecutionToolResultError

```yaml
{"example": {"error_code": "unavailable", "type": "bash_code_execution_tool_result_error"}, "properties": {"error_code": {"enum": ["invalid_tool_input", "unavailable", "too_many_requests", "execution_time_exceeded", "output_file_too_large"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["bash_code_execution_tool_result_error"], "type": "string"}}, "required": ["error_code", "type"], "type": "object"}
```
