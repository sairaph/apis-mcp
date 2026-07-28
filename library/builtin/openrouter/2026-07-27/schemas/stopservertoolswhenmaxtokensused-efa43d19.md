---
title: StopServerToolsWhenMaxTokensUsed
page_id: schema-stopservertoolswhenmaxtokensused-efa43d19
path: schemas
description: Stop once cumulative token usage across the loop exceeds this threshold.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenMaxTokensUsed

Stop once cumulative token usage across the loop exceeds this threshold.

```yaml
{"description": "Stop once cumulative token usage across the loop exceeds this threshold.", "example": {"max_tokens": 10000, "type": "max_tokens_used"}, "properties": {"max_tokens": {"type": "integer"}, "type": {"enum": ["max_tokens_used"], "type": "string"}}, "required": ["type", "max_tokens"], "type": "object"}
```
