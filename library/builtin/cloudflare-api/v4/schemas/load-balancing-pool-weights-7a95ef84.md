---
title: load-balancing_pool_weights
page_id: schema-load-balancing-pool-weights-7a95ef84
path: schemas
description: A mapping of pool IDs to custom weights, relative to the other pools. The declarative alternative to `random_steering.pool_weights`; mutually exclusive with `random_steering`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool_weights

A mapping of pool IDs to custom weights, relative to the other pools. The declarative alternative to `random_steering.pool_weights`; mutually exclusive with `random_steering`.

```yaml
{"description": "A mapping of pool IDs to custom weights, relative to the other pools. The declarative alternative to `random_steering.pool_weights`; mutually exclusive with `random_steering`.", "type": "object", "example": {"9290f38c5d07c2e2f4df57b1f61d4196": 0.5, "de90f38ced07c2e2f4df50b1f61d4194": 0.3}, "additionalProperties": {"type": "number", "x-auditable": true}, "x-stainless-terraform-configurability": "computed_optional"}
```
