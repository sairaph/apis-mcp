---
title: ChatCompletionResponse
page_id: schema-chatcompletionresponse-16975903
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ChatCompletionResponse

```yaml
{"type": "object", "properties": {"id": {"description": "Task ID", "type": "string"}, "request_id": {"description": "Request ID", "type": "string"}, "created": {"description": "Request creation time, Unix timestamp in seconds", "type": "integer"}, "model": {"description": "Model name", "type": "string"}, "choices": {"type": "array", "description": "List of model responses", "items": {"type": "object", "properties": {"index": {"type": "integer", "description": "Result index."}, "message": {"$ref": "#/components/schemas/ChatCompletionResponseMessage"}, "finish_reason": {"type": "string", "description": "Reason for model inference termination. Can be `stop`, `tool_calls`, `length`, `sensitive`, `model_context_window_exceeded` or `network_error`."}}}}, "usage": {"type": "object", "description": "Token usage statistics returned when the model call ends.", "properties": {"prompt_tokens": {"type": "number", "description": "Number of tokens in user input"}, "completion_tokens": {"type": "number", "description": "Number of output tokens"}, "prompt_tokens_details": {"type": "object", "properties": {"cached_tokens": {"type": "number", "description": "Number of tokens served from cache"}}}, "total_tokens": {"type": "integer", "description": "Total number of tokens"}}}, "web_search": {"description": "Search results.", "type": "array", "items": {"$ref": "#/components/schemas/WebSearchObjectResponse"}}}}
```
