---
title: zero-trust-gateway_type
page_id: schema-zero-trust-gateway-type-230e43ba
path: schemas
description: Indicate the read-only certificate type, BYO-PKI (custom) or Gateway-managed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_type

Indicate the read-only certificate type, BYO-PKI (custom) or Gateway-managed.

```yaml
{"description": "Indicate the read-only certificate type, BYO-PKI (custom) or Gateway-managed.", "type": "string", "example": "gateway_managed", "enum": ["custom", "gateway_managed"], "readOnly": true, "x-auditable": true}
```
