---
title: MessagesContentBlockStartEvent
page_id: schema-messagescontentblockstartevent-84acf28b
path: schemas
description: Event sent when a new content block starts
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesContentBlockStartEvent

Event sent when a new content block starts

```yaml
{"description": "Event sent when a new content block starts", "example": {"content_block": {"citations": [], "text": "", "type": "text"}, "index": 0, "type": "content_block_start"}, "properties": {"content_block": {"anyOf": [{"$ref": "#/components/schemas/AnthropicTextBlock"}, {"$ref": "#/components/schemas/AnthropicToolUseBlock"}, {"$ref": "#/components/schemas/AnthropicThinkingBlock"}, {"$ref": "#/components/schemas/AnthropicRedactedThinkingBlock"}, {"$ref": "#/components/schemas/ORAnthropicServerToolUseBlock"}, {"$ref": "#/components/schemas/AnthropicWebSearchToolResult"}, {"$ref": "#/components/schemas/AnthropicWebFetchToolResult"}, {"$ref": "#/components/schemas/AnthropicCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicBashCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicToolSearchToolResult"}, {"$ref": "#/components/schemas/AnthropicContainerUpload"}, {"$ref": "#/components/schemas/AnthropicCompactionBlock"}, {"$ref": "#/components/schemas/AnthropicAdvisorToolResult"}, {"properties": {"content": {"type": ["string", "null"]}, "encrypted_content": {"type": ["string", "null"]}, "type": {"enum": ["compaction"], "type": "string"}}, "required": ["type", "content"], "type": "object"}]}, "index": {"type": "integer"}, "type": {"enum": ["content_block_start"], "type": "string"}}, "required": ["type", "index", "content_block"], "type": "object"}
```
