---
title: quotes_resource_computed
page_id: schema-quotes-resource-computed-b5866e3b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_computed

```yaml
{"title": "QuotesResourceComputed", "required": ["upfront"], "type": "object", "properties": {"recurring": {"description": "The definitive totals and line items the customer will be charged on a recurring basis. Takes into account the line items with recurring prices and discounts with `duration=forever` coupons only. Defaults to `null` if no inputted line items with recurring prices.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/quotes_resource_recurring"}]}, "upfront": {"$ref": "#/components/schemas/quotes_resource_upfront"}}, "description": "", "x-expandableFields": ["recurring", "upfront"]}
```
