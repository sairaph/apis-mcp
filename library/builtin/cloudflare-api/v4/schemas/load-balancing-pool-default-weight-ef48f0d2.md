---
title: load-balancing_pool_default_weight
page_id: schema-load-balancing-pool-default-weight-ef48f0d2
path: schemas
description: The default weight for pools not listed in `pool_weights`. The declarative alternative to `random_steering.default_weight`; mutually exclusive with `random_steering`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool_default_weight

The default weight for pools not listed in `pool_weights`. The declarative alternative to `random_steering.default_weight`; mutually exclusive with `random_steering`.

```yaml
{"description": "The default weight for pools not listed in `pool_weights`. The declarative alternative to `random_steering.default_weight`; mutually exclusive with `random_steering`.", "type": "number", "example": 0.2, "maximum": 1, "minimum": 0, "multipleOf": 0.1, "x-stainless-terraform-configurability": "computed_optional"}
```
