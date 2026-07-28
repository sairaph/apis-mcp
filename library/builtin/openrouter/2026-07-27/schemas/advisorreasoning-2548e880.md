---
title: AdvisorReasoning
page_id: schema-advisorreasoning-2548e880
path: schemas
description: Reasoning configuration forwarded to the advisor call. Use this to control reasoning effort and token budget for models that support extended thinking.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AdvisorReasoning

Reasoning configuration forwarded to the advisor call. Use this to control reasoning effort and token budget for models that support extended thinking.

```yaml
{"description": "Reasoning configuration forwarded to the advisor call. Use this to control reasoning effort and token budget for models that support extended thinking.", "example": {"effort": "high"}, "properties": {"effort": {"description": "Reasoning effort level for the advisor call.", "enum": ["max", "xhigh", "high", "medium", "low", "minimal", "none"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "max_tokens": {"description": "Maximum number of reasoning tokens the advisor may use.", "type": "integer"}}, "type": "object"}
```
