---
title: AnthropicAdvisorToolResult
page_id: schema-anthropicadvisortoolresult-858580d0
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicAdvisorToolResult

```yaml
{"example": {"content": {"text": "Advisor response text", "type": "advisor_result"}, "tool_use_id": "srvtoolu_01abc", "type": "advisor_tool_result"}, "properties": {"content": {"additionalProperties": {}, "type": "object"}, "tool_use_id": {"type": "string"}, "type": {"enum": ["advisor_tool_result"], "type": "string"}}, "required": ["type", "tool_use_id", "content"], "type": "object"}
```
