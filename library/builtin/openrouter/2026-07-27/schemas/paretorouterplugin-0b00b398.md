---
title: ParetoRouterPlugin
page_id: schema-paretorouterplugin-0b00b398
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ParetoRouterPlugin

```yaml
{"example": {"enabled": true, "id": "pareto-router", "max_price": 5, "price_source": "prompt"}, "properties": {"enabled": {"description": "Set to false to disable the pareto-router plugin for this request. Defaults to true.", "type": "boolean"}, "id": {"enum": ["pareto-router"], "type": "string"}, "max_price": {"description": "Maximum input price in USD per million tokens. When set, quality-tier selection (min_coding_score) is bypassed: the router computes the Pareto frontier over the top coding models and routes to the best-scoring frontier model priced at or below this cap, falling back through cheaper frontier models, then non-frontier models. Enforced against the price source given by price_source. Returns 404 when no candidate satisfies the cap.", "example": 5, "format": "double", "minimum": 0, "type": "number"}, "min_coding_score": {"description": "Minimum coding quality score between 0 and 1. Maps to internal quality tiers: >= 0.66 → high (top coding models), >= 0.33 → medium (strong modern flagships), < 0.33 → low (capable coders above the median). Omit to default to the highest tier (equivalent to >= 0.66). Not used when max_price is set (price-based selection takes over).", "example": 0.8, "format": "double", "maximum": 1, "minimum": 0, "type": "number"}, "price_source": {"description": "Price source for the Pareto frontier cost axis and for enforcing max_price. \"prompt\" uses catalog list price (endpoint.pricing.prompt). \"weighted_avg\" uses traffic-weighted effective input price from ClickHouse, falling back to prompt price for models without traffic data. Defaults to \"prompt\".", "enum": ["prompt", "weighted_avg"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["id"], "type": "object"}
```
