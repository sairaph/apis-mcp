---
title: aig-billing_GetTopupLimitsResult
page_id: schema-aig-billing-gettopuplimitsresult-759f3721
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetTopupLimitsResult

```yaml
{"type": "object", "properties": {"currency": {"description": "ISO 4217 currency code.", "type": "string", "example": "USD"}, "max_cents": {"type": "integer"}, "min_cents": {"description": "Minimum allowed top-up amount in cents.", "type": "integer", "example": 1000}}, "required": ["min_cents", "max_cents", "currency"]}
```
