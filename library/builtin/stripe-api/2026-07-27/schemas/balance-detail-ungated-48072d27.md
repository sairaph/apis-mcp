---
title: balance_detail_ungated
page_id: schema-balance-detail-ungated-48072d27
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_detail_ungated

```yaml
{"title": "BalanceDetailUngated", "required": ["available", "pending"], "type": "object", "properties": {"available": {"type": "array", "description": "Funds that are available for use.", "items": {"$ref": "#/components/schemas/balance_amount"}}, "pending": {"type": "array", "description": "Funds that are pending", "items": {"$ref": "#/components/schemas/balance_amount"}}}, "description": "", "x-expandableFields": ["available", "pending"]}
```
