---
title: CommonAgentResultResponse
page_id: schema-commonagentresultresponse-67d7eb6f
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# CommonAgentResultResponse

```yaml
{"type": "object", "properties": {"status": {"type": "string", "description": "`pending` (processing), `success` (completed), `failed` (failed)."}, "agent_id": {"type": "string", "description": "Agent ID"}, "async_id": {"type": "string", "description": "Asynchronous task ID."}, "choices": {"type": "array", "description": "Agent output.", "items": {"type": "object", "properties": {"index": {"type": "integer", "description": "Result index."}, "finish_reason": {"type": "string", "description": "Reason for model inference termination. Can be ‘stop’, ‘tool_calls’, ‘length’, ‘sensitive’, or ‘network_error’."}, "message": {"type": "array", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "Role: fixed as `assistant`."}, "content": {"type": "array", "description": "Video file metadata", "items": {"type": "object", "properties": {"type": {"type": "string", "description": "object type: `video_url`."}, "video_url": {"type": "string", "description": "MP4 video URL."}}}}}}}}}}}}
```
