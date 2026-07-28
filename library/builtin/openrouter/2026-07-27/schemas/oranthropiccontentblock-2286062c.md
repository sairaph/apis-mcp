---
title: ORAnthropicContentBlock
page_id: schema-oranthropiccontentblock-2286062c
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ORAnthropicContentBlock

```yaml
{"discriminator": {"mapping": {"advisor_tool_result": "#/components/schemas/AnthropicAdvisorToolResult", "bash_code_execution_tool_result": "#/components/schemas/AnthropicBashCodeExecutionToolResult", "code_execution_tool_result": "#/components/schemas/AnthropicCodeExecutionToolResult", "compaction": "#/components/schemas/AnthropicCompactionBlock", "container_upload": "#/components/schemas/AnthropicContainerUpload", "redacted_thinking": "#/components/schemas/AnthropicRedactedThinkingBlock", "server_tool_use": "#/components/schemas/ORAnthropicServerToolUseBlock", "text": "#/components/schemas/AnthropicTextBlock", "text_editor_code_execution_tool_result": "#/components/schemas/AnthropicTextEditorCodeExecutionToolResult", "thinking": "#/components/schemas/AnthropicThinkingBlock", "tool_search_tool_result": "#/components/schemas/AnthropicToolSearchToolResult", "tool_use": "#/components/schemas/AnthropicToolUseBlock", "web_fetch_tool_result": "#/components/schemas/AnthropicWebFetchToolResult", "web_search_tool_result": "#/components/schemas/AnthropicWebSearchToolResult"}, "propertyName": "type"}, "example": {"citations": null, "text": "Hello, world!", "type": "text"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicTextBlock"}, {"$ref": "#/components/schemas/AnthropicToolUseBlock"}, {"$ref": "#/components/schemas/AnthropicThinkingBlock"}, {"$ref": "#/components/schemas/AnthropicRedactedThinkingBlock"}, {"$ref": "#/components/schemas/ORAnthropicServerToolUseBlock"}, {"$ref": "#/components/schemas/AnthropicWebSearchToolResult"}, {"$ref": "#/components/schemas/AnthropicWebFetchToolResult"}, {"$ref": "#/components/schemas/AnthropicCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicBashCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicTextEditorCodeExecutionToolResult"}, {"$ref": "#/components/schemas/AnthropicToolSearchToolResult"}, {"$ref": "#/components/schemas/AnthropicContainerUpload"}, {"$ref": "#/components/schemas/AnthropicCompactionBlock"}, {"$ref": "#/components/schemas/AnthropicAdvisorToolResult"}]}
```
