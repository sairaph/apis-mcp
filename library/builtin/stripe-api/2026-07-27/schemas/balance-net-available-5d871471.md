---
title: balance_net_available
page_id: schema-balance-net-available-5d871471
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_net_available

```yaml
{"title": "BalanceNetAvailable", "required": ["amount", "destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Net balance amount, subtracting fees from platform-set pricing."}, "destination": {"maxLength": 5000, "type": "string", "description": "ID of the external account for this net balance (not expandable)."}, "source_types": {"$ref": "#/components/schemas/balance_amount_by_source_type"}}, "description": "", "x-expandableFields": ["source_types"]}
```
