---
title: DABenchmarkEntry
page_id: schema-dabenchmarkentry-4cc247d5
path: schemas
description: A single Design Arena benchmark entry for a specific arena+category
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# DABenchmarkEntry

A single Design Arena benchmark entry for a specific arena+category

```yaml
{"description": "A single Design Arena benchmark entry for a specific arena+category", "example": {"arena": "models", "category": "website", "elo": 1385.2, "rank": 5, "win_rate": 62.5}, "properties": {"arena": {"description": "Arena type (e.g. models, builders, agents)", "example": "models", "type": "string"}, "category": {"description": "Category within the arena (e.g. website, gamedev, uicomponent)", "example": "website", "type": "string"}, "elo": {"description": "ELO rating from head-to-head arena battles", "example": 1385.2, "format": "double", "type": "number"}, "rank": {"description": "Rank position within this arena+category among models available on OpenRouter (1 = highest ELO)", "example": 5, "type": "integer"}, "win_rate": {"description": "Win rate percentage in arena battles", "example": 62.5, "format": "double", "type": "number"}}, "required": ["arena", "category", "elo", "win_rate", "rank"], "type": "object"}
```
