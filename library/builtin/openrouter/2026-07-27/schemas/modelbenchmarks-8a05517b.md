---
title: ModelBenchmarks
page_id: schema-modelbenchmarks-8a05517b
path: schemas
description: Third-party benchmark rankings for this model. Omitted when no benchmark data is available.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ModelBenchmarks

Third-party benchmark rankings for this model. Omitted when no benchmark data is available.

```yaml
{"description": "Third-party benchmark rankings for this model. Omitted when no benchmark data is available.", "example": {"artificial_analysis": {"agentic_index": 55.8, "coding_index": 63.2, "intelligence_index": 71.4}, "design_arena": [{"arena": "models", "category": "website", "elo": 1385.2, "rank": 5, "win_rate": 62.5}]}, "properties": {"artificial_analysis": {"$ref": "#/components/schemas/AABenchmarkEntry"}, "design_arena": {"description": "Design Arena ELO rankings across arena+category pairs.", "example": [{"arena": "models", "category": "website", "elo": 1385.2, "rank": 5, "win_rate": 62.5}], "items": {"$ref": "#/components/schemas/DABenchmarkEntry"}, "type": "array"}}, "required": ["design_arena"], "type": "object"}
```
