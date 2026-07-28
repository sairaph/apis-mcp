---
title: load-balancing_country_pools
page_id: schema-load-balancing-country-pools-9893dde8
path: schemas
description: A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_country_pools

A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.

```yaml
{"description": "A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.", "type": "object", "example": {"GB": ["abd90f38ced07c2e2f4df50b1f61d4194"], "US": ["de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"]}, "additionalProperties": {"description": "A `string:[string]` object of key-values. Country code maps to list of pool IDs.", "items": {"type": "string", "x-auditable": true}, "type": "array"}}
```
