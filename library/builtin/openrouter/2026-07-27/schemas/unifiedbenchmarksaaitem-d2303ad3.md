---
title: UnifiedBenchmarksAAItem
page_id: schema-unifiedbenchmarksaaitem-d2303ad3
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnifiedBenchmarksAAItem

```yaml
{"example": {"agentic_index": 58.3, "coding_index": 65.8, "display_name": "GPT-4o", "intelligence_index": 71.2, "model_permaslug": "openai/gpt-4o", "pricing": {"completion": "0.00001", "prompt": "0.0000025"}, "source": "artificial-analysis"}, "properties": {"agentic_index": {"description": "Artificial Analysis Agentic Index composite score. Higher is better.", "example": 58.3, "format": "double", "type": ["number", "null"]}, "coding_index": {"description": "Artificial Analysis Coding Index composite score. Higher is better.", "example": 65.8, "format": "double", "type": ["number", "null"]}, "display_name": {"description": "Model name as listed on Artificial Analysis.", "example": "GPT-4o", "type": "string"}, "intelligence_index": {"description": "Artificial Analysis Intelligence Index composite score. Higher is better.", "example": 71.2, "format": "double", "type": ["number", "null"]}, "model_permaslug": {"description": "Stable OpenRouter model identifier.", "example": "openai/gpt-4o", "type": "string"}, "pricing": {"$ref": "#/components/schemas/UnifiedBenchmarkPricing"}, "source": {"description": "Benchmark source discriminator.", "enum": ["artificial-analysis"], "type": "string"}}, "required": ["source", "model_permaslug", "display_name", "intelligence_index", "coding_index", "agentic_index", "pricing"], "type": "object"}
```
