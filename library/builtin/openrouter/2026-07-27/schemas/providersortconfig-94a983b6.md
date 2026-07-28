---
title: ProviderSortConfig
page_id: schema-providersortconfig-94a983b6
path: schemas
description: The provider sorting strategy (price, throughput, latency)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ProviderSortConfig

The provider sorting strategy (price, throughput, latency)

```yaml
{"description": "The provider sorting strategy (price, throughput, latency)", "example": {"by": "price", "partition": "model"}, "properties": {"by": {"description": "The provider sorting strategy (price, throughput, latency)", "enum": ["price", "throughput", "latency", "exacto", null], "example": "price", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "partition": {"description": "Partitioning strategy for sorting: \"model\" (default) groups endpoints by model before sorting (fallback models remain fallbacks), \"none\" sorts all endpoints together regardless of model.", "enum": ["model", "none", null], "example": "model", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}}, "type": "object"}
```
