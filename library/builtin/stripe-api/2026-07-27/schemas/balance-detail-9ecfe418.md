---
title: balance_detail
page_id: schema-balance-detail-9ecfe418
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_detail

```yaml
{"title": "BalanceDetail", "required": ["available"], "type": "object", "properties": {"available": {"type": "array", "description": "Funds that are available for use.", "items": {"$ref": "#/components/schemas/balance_amount"}}}, "description": "", "x-expandableFields": ["available"]}
```
