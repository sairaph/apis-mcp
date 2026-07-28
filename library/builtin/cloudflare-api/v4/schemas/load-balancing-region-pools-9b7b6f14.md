---
title: load-balancing_region_pools
page_id: schema-load-balancing-region-pools-9b7b6f14
path: schemas
description: A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_region_pools

A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.

```yaml
{"description": "A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.", "type": "object", "example": {"ENAM": ["00920f38ce07c2e2f4df50b1f61d4194"], "WNAM": ["de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"]}, "additionalProperties": {"description": "A `string:[string]` object of key-values. Region code maps to list of pool IDs.", "items": {"type": "string", "x-auditable": true}, "type": "array"}}
```
