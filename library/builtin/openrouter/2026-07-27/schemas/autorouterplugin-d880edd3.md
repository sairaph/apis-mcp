---
title: AutoRouterPlugin
page_id: schema-autorouterplugin-d880edd3
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AutoRouterPlugin

```yaml
{"example": {"allowed_models": ["anthropic/*", "openai/gpt-4o"], "cost_quality_tradeoff": 7, "enabled": true, "id": "auto-router", "pin_model": false}, "properties": {"allowed_models": {"description": "List of model patterns to filter which models the auto-router can route between. Supports wildcards (e.g., \"anthropic/*\" matches all Anthropic models). When not specified, uses the default supported models list.", "example": ["anthropic/*", "openai/gpt-4o", "google/*"], "items": {"type": "string"}, "type": "array"}, "cost_quality_tradeoff": {"description": "Controls cost vs. quality routing tradeoff (0–10). 0 = pure quality (best model regardless of cost), 10 = maximize for cost (cheapest model wins). Intermediate values blend quality and cost signals continuously. Defaults to 7.", "example": 7, "maximum": 10, "minimum": 0, "type": "integer"}, "enabled": {"description": "Set to false to disable the auto-router plugin for this request. Defaults to true.", "type": "boolean"}, "id": {"enum": ["auto-router"], "type": "string"}, "pin_model": {"description": "When true, reuses the model from the most recent assistant message's `model` attribute for subsequent turns. Defaults to false.", "example": false, "type": "boolean"}}, "required": ["id"], "type": "object"}
```
