---
title: load-balancing_default_pools
page_id: schema-load-balancing-default-pools-b36986ea
path: schemas
description: A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_default_pools

A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.

```yaml
{"description": "A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.", "type": "array", "items": {"description": "A pool ID.", "type": "string"}, "example": ["17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"]}
```
