---
title: zero-trust-gateway_filters
page_id: schema-zero-trust-gateway-filters-9d854abe
path: schemas
description: Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions. Can only contain a single value.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_filters

Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions. Can only contain a single value.

```yaml
{"description": "Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions. Can only contain a single value.", "type": "array", "items": {"description": "Specify the protocol or layer to use.", "enum": ["http", "dns", "l4", "egress", "dns_resolver"], "example": "http", "type": "string", "x-auditable": true}, "example": ["http"], "maxItems": 1, "minItems": 1}
```
