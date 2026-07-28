---
title: FusionAnalysisResult
page_id: schema-fusionanalysisresult-884cf312
path: schemas
description: Structured analysis produced by the fusion judge model.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionAnalysisResult

Structured analysis produced by the fusion judge model.

```yaml
{"description": "Structured analysis produced by the fusion judge model.", "example": {"blind_spots": ["No model considered the impact on existing API consumers."], "consensus": ["All panel models agree the request is asking for a concise summary."], "contradictions": [{"stances": [{"model": "openai/gpt-5", "stance": "Favors an incremental rollout."}, {"model": "anthropic/claude-sonnet-4.5", "stance": "Favors a single coordinated migration."}], "topic": "Recommended approach"}], "partial_coverage": [{"models": ["openai/gpt-5"], "point": "Only one model addressed the rollback strategy."}], "unique_insights": [{"insight": "Highlighted a backwards-compatibility risk the other models missed.", "model": "anthropic/claude-sonnet-4.5"}]}, "properties": {"blind_spots": {"items": {"type": "string"}, "type": "array"}, "consensus": {"items": {"type": "string"}, "type": "array"}, "contradictions": {"items": {"properties": {"stances": {"items": {"properties": {"model": {"type": "string"}, "stance": {"type": "string"}}, "required": ["model", "stance"], "type": "object"}, "type": "array"}, "topic": {"type": "string"}}, "required": ["topic", "stances"], "type": "object"}, "type": "array"}, "partial_coverage": {"items": {"properties": {"models": {"items": {"type": "string"}, "type": "array"}, "point": {"type": "string"}}, "required": ["models", "point"], "type": "object"}, "type": "array"}, "unique_insights": {"items": {"properties": {"insight": {"type": "string"}, "model": {"type": "string"}}, "required": ["model", "insight"], "type": "object"}, "type": "array"}}, "required": ["consensus", "contradictions", "partial_coverage", "unique_insights", "blind_spots"], "type": "object"}
```
