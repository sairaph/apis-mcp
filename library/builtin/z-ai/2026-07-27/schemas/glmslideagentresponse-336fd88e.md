---
title: GlmSlideAgentResponse
page_id: schema-glmslideagentresponse-336fd88e
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# GlmSlideAgentResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Request ID"}, "conversation_id": {"type": "string", "description": "Conversation ID"}, "agent_id": {"type": "string", "description": "Agent ID"}, "choices": {"type": "array", "description": "Agent output.", "items": {"type": "object", "properties": {"index": {"type": "integer", "description": "Result index."}, "finish_reason": {"type": "string", "description": "Reason for model inference termination. Can be ‘stop’, ‘tool_calls’, ‘length’, ‘sensitive’, or ‘network_error’."}, "message": {"type": "array", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "Role: fixed as `assistant`."}, "phase": {"type": "string", "description": "Current role type: thinking、tool、answer"}, "content": {"type": "array", "description": "Content metadata", "items": {"type": "object", "properties": {"type": {"type": "string", "description": "Response Content type: text、object"}, "tag_cn": {"type": "string", "description": "CN Tag."}, "tag_en": {"type": "string", "description": "EN Tag."}, "text": {"type": "string", "description": "Output string content when type is text"}, "object": {"type": "object", "description": "Output object content when type is object", "properties": {"tool_name": {"type": "string", "description": "Tool name eg: search、insert_page"}, "input": {"type": "string", "description": "Tool input content"}, "output": {"type": "string", "description": "Tool output content, will output html when generate slide"}, "position": {"type": "array", "description": "If the tool involves operations on a PPT file, the position field specifies which slides are being manipulated.\n If the user says, “Insert a slide after the second slide,” then position = [3], and the output is the HTML content of the third slide.\n If the user says, “Please delete slides 4, 5, and 6,” then position = [4, 5, 6].", "items": {"type": "number"}}}}}}}}}}}}}, "error": {"type": "object", "properties": {"code": {"type": "string", "description": "Error code."}, "message": {"type": "string", "description": "Error message."}}}}}
```
