---
title: UnifiedBenchmarksDAItem
page_id: schema-unifiedbenchmarksdaitem-a8889bf1
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnifiedBenchmarksDAItem

```yaml
{"example": {"arena": "models", "avg_generation_time_ms": 3200, "category": "codecategories", "display_name": "Claude Sonnet 4", "elo": 1423, "model_permaslug": "anthropic/claude-sonnet-4", "pricing": {"completion": "0.000015", "prompt": "0.000003"}, "source": "design-arena", "tournament_stats": {"first_place": 12, "fourth_place": 2, "second_place": 8, "third_place": 5, "total": 27}, "win_rate": 72}, "properties": {"arena": {"description": "Arena this ranking belongs to.", "example": "models", "type": "string"}, "avg_generation_time_ms": {"description": "Average generation time in milliseconds.", "example": 3200, "format": "double", "type": ["number", "null"]}, "category": {"description": "Category within the arena.", "example": "codecategories", "type": "string"}, "display_name": {"description": "Human-readable model name from Design Arena.", "example": "Claude Sonnet 4", "type": "string"}, "elo": {"description": "ELO rating from head-to-head arena battles.", "example": 1423, "format": "double", "type": "number"}, "model_permaslug": {"description": "Stable OpenRouter model identifier when mapped; otherwise the upstream Design Arena model id.", "example": "anthropic/claude-sonnet-4", "type": "string"}, "pricing": {"$ref": "#/components/schemas/UnifiedBenchmarkPricing"}, "source": {"description": "Benchmark source discriminator.", "enum": ["design-arena"], "type": "string"}, "tournament_stats": {"description": "Placement distribution from tournament matches.", "properties": {"first_place": {"type": ["integer", "null"]}, "fourth_place": {"type": ["integer", "null"]}, "second_place": {"type": ["integer", "null"]}, "third_place": {"type": ["integer", "null"]}, "total": {"type": ["integer", "null"]}}, "required": ["first_place", "second_place", "third_place", "fourth_place", "total"], "type": "object"}, "win_rate": {"description": "Win rate as a percentage (0–100).", "example": 72, "format": "double", "type": "number"}}, "required": ["source", "model_permaslug", "display_name", "arena", "category", "elo", "win_rate", "avg_generation_time_ms", "tournament_stats", "pricing"], "type": "object"}
```
