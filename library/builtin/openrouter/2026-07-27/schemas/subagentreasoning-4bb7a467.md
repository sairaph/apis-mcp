---
title: SubagentReasoning
page_id: schema-subagentreasoning-4bb7a467
path: schemas
description: Reasoning configuration forwarded to the subagent call. Use this to control reasoning effort and token budget for models that support extended thinking.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SubagentReasoning

Reasoning configuration forwarded to the subagent call. Use this to control reasoning effort and token budget for models that support extended thinking.

```yaml
{"description": "Reasoning configuration forwarded to the subagent call. Use this to control reasoning effort and token budget for models that support extended thinking.", "example": {"effort": "low"}, "properties": {"effort": {"description": "Reasoning effort level for the subagent call.", "enum": ["max", "xhigh", "high", "medium", "low", "minimal", "none"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "max_tokens": {"description": "Maximum number of reasoning tokens the subagent may use. Accepted and validated but not yet forwarded to the subagent call.", "type": "integer"}}, "type": "object"}
```
