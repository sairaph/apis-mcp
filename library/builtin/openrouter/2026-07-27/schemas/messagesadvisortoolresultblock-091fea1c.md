---
title: MessagesAdvisorToolResultBlock
page_id: schema-messagesadvisortoolresultblock-091fea1c
path: schemas
description: Advisor tool result from a prior assistant turn, replayed back to the model on the next turn. Mirrors the block Anthropic returns in assistant content when the `advisor_20260301` tool runs.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesAdvisorToolResultBlock

Advisor tool result from a prior assistant turn, replayed back to the model on the next turn. Mirrors the block Anthropic returns in assistant content when the `advisor_20260301` tool runs.

```yaml
{"description": "Advisor tool result from a prior assistant turn, replayed back to the model on the next turn. Mirrors the block Anthropic returns in assistant content when the `advisor_20260301` tool runs.", "example": {"content": {"text": "Advisor response text", "type": "advisor_result"}, "tool_use_id": "srvtoolu_01abc", "type": "advisor_tool_result"}, "properties": {"content": {"additionalProperties": {}, "type": "object"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["advisor_tool_result"], "type": "string"}}, "required": ["type", "tool_use_id", "content"], "type": "object"}
```
