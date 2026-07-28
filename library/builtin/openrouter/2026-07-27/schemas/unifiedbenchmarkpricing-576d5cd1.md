---
title: UnifiedBenchmarkPricing
page_id: schema-unifiedbenchmarkpricing-576d5cd1
path: schemas
description: OpenRouter pricing per token for this model. Null if pricing is unavailable.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnifiedBenchmarkPricing

OpenRouter pricing per token for this model. Null if pricing is unavailable.

```yaml
{"description": "OpenRouter pricing per token for this model. Null if pricing is unavailable.", "example": {"completion": "0.000015", "prompt": "0.000003"}, "properties": {"completion": {"description": "Cost per output token (USD, decimal string).", "example": "0.000015", "type": "string"}, "prompt": {"description": "Cost per input token (USD, decimal string).", "example": "0.000003", "type": "string"}}, "required": ["prompt", "completion"], "type": ["object", "null"]}
```
