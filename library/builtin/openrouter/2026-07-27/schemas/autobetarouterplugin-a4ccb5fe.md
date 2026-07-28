---
title: AutoBetaRouterPlugin
page_id: schema-autobetarouterplugin-a4ccb5fe
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AutoBetaRouterPlugin

```yaml
{"example": {"allowed_models": ["anthropic/*", "openai/gpt-4o"], "cost_quality_tradeoff": 9, "enabled": true, "id": "auto-beta-router"}, "properties": {"allowed_models": {"description": "List of model patterns to filter which models the auto-beta-router can route between. Supports wildcards (e.g., \"anthropic/*\" matches all Anthropic models). When not specified, uses the default supported models list.", "example": ["anthropic/*", "openai/gpt-4o", "google/*"], "items": {"type": "string"}, "type": "array"}, "cost_quality_tradeoff": {"description": "Balances routing between cost and quality on a 0-10 scale. The auto-beta-router ranks models for the classified task type by community spend share, then filters candidates by their average cost per generation for that task. Higher values favor cheaper models: 10 keeps only models around the cheapest 10th percentile, while 0 permits models up to the 90th percentile for cost. Defaults to 9.", "example": 9, "maximum": 10, "minimum": 0, "type": "integer"}, "enabled": {"description": "Set to false to disable the auto-beta-router plugin for this request. Defaults to true.", "type": "boolean"}, "id": {"enum": ["auto-beta-router"], "type": "string"}}, "required": ["id"], "type": "object"}
```
