---
title: TokenizerRequest
page_id: schema-tokenizerrequest-c45d62da
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# TokenizerRequest

```yaml
{"type": "object", "required": ["model", "messages"], "properties": {"model": {"type": "string", "description": "The model code to be called.", "example": "glm-4.6", "default": "glm-4.6", "enum": ["glm-4.6", "glm-4.6v", "glm-4.5"]}, "messages": {"type": "array", "description": "The current conversation message list as the model’s prompt input, provided in JSON array format, e.g.,`{“role”: “user”, “content”: “Hello”}`. Possible message types include system messages, user messages. Note: The input must not consist of system or assistant messages only.", "items": {"oneOf": [{"title": "User Message", "type": "object", "properties": {"role": {"type": "string", "enum": ["user"], "description": "Role of the message author", "default": "user"}, "content": {"oneOf": [{"type": "array", "description": "Multimodal message content, supports text, images, video, file", "items": {"$ref": "#/components/schemas/VisionMultimodalContentItem"}}, {"type": "string", "description": "Text message content (can switch to multimodal message above)", "example": "What opportunities and challenges will the Chinese large model industry face in 2025?"}]}}, "required": ["role", "content"]}, {"title": "System Message", "type": "object", "properties": {"role": {"type": "string", "enum": ["system"], "description": "Role of the message author", "default": "system"}, "content": {"oneOf": [{"type": "string", "description": "Message text content", "example": "You are a helpful assistant."}]}}, "required": ["role", "content"]}, {"title": "Assistant Message", "type": "object", "description": "Can include tool calls", "properties": {"role": {"type": "string", "enum": ["assistant"], "description": "Role of the message author", "default": "assistant"}, "content": {"oneOf": [{"type": "string", "description": "Text message content", "example": "I'll help you with that analysis."}]}}, "required": ["role"]}]}, "minItems": 1}, "tools": {"type": "array", "description": "List of tools the model can call. Supports up to `128` functions.", "anyOf": [{"items": {"$ref": "#/components/schemas/FunctionToolSchema"}}]}, "request_id": {"type": "string", "description": "Passed by the user side, needs to be unique; used to distinguish each request, 6–64 characters. If not provided by the user side, the platform will generate one by default.", "minLength": 6, "maxLength": 64}, "user_id": {"type": "string", "description": "Unique ID for the end user, 6–128 characters. Avoid using sensitive information.", "minLength": 6, "maxLength": 128}}}
```
