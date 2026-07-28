---
title: TranslationAgentResponse
page_id: schema-translationagentresponse-e94f2fd6
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# TranslationAgentResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Task ID."}, "agent_id": {"type": "string", "description": "Agent ID."}, "status": {"type": "string", "description": "Task status."}, "choices": {"type": "array", "description": "Model output content.", "items": {"type": "object", "properties": {"index": {"type": "integer", "description": "Result index."}, "finish_reason": {"type": "string", "description": "Termination reason: `stop` (normal completion), `tool_calls` (model calls), `length` (token limit exceeded), `sensitive` (content flagged), `network_error` (model inference error)."}, "messages": {"type": "object", "description": "Model response message.", "properties": {"role": {"type": "string", "description": "Dialog role (default: `assistant`)."}, "content": {"type": "object", "description": "Inference result", "properties": {"type": {"type": "string", "description": "Result type."}, "text": {"type": "string", "description": "Result content."}}}}}}}}, "usage": {"type": "object", "description": "Token usage statistics.", "properties": {"prompt_tokens": {"type": "integer", "description": "Input tokens count."}, "completion_tokens": {"type": "integer", "description": "Output tokens count."}, "total_tokens": {"type": "integer", "description": "Total tokens count."}, "total_calls": {"type": "integer", "description": "Total number of calls"}}}}}
```
