---
title: UnifiedBenchmarksMeta
page_id: schema-unifiedbenchmarksmeta-95bf7c34
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnifiedBenchmarksMeta

```yaml
{"example": {"as_of": "2026-06-03T12:00:00Z", "citation": "Source: Artificial Analysis (artificialanalysis.ai) via OpenRouter (openrouter.ai/rankings).", "model_count": 50, "source": "artificial-analysis", "source_url": "https://artificialanalysis.ai", "task_type": null, "version": "v1"}, "properties": {"as_of": {"description": "ISO-8601 timestamp of when this data was last updated.", "example": "2026-06-03T12:00:00Z", "type": "string"}, "citation": {"description": "Required attribution when republishing this data, or null when results span multiple sources (attribute each item individually by its `source` discriminator).", "example": "Source: Artificial Analysis (artificialanalysis.ai) via OpenRouter (openrouter.ai/rankings).", "type": ["string", "null"]}, "model_count": {"description": "Number of unique models in the response.", "type": "integer"}, "source": {"description": "The source filter applied, or null when all sources are returned.", "enum": ["artificial-analysis", "design-arena", null], "example": "artificial-analysis", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "source_url": {"description": "URL of the upstream data source, or null when results span multiple sources.", "example": "https://artificialanalysis.ai", "type": ["string", "null"]}, "task_type": {"description": "The task_type filter applied, or null if showing all.", "type": ["string", "null"]}, "version": {"description": "Dataset version.", "enum": ["v1"], "type": "string"}}, "required": ["as_of", "version", "source", "source_url", "citation", "model_count", "task_type"], "type": "object"}
```
