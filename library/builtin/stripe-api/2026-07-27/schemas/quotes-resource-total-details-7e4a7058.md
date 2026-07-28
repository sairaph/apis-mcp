---
title: quotes_resource_total_details
page_id: schema-quotes-resource-total-details-7e4a7058
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_total_details

```yaml
{"title": "QuotesResourceTotalDetails", "required": ["amount_discount", "amount_tax"], "type": "object", "properties": {"amount_discount": {"type": "integer", "description": "This is the sum of all the discounts."}, "amount_shipping": {"type": "integer", "description": "This is the sum of all the shipping amounts.", "nullable": true}, "amount_tax": {"type": "integer", "description": "This is the sum of all the tax amounts."}, "breakdown": {"$ref": "#/components/schemas/quotes_resource_total_details_resource_breakdown"}}, "description": "", "x-expandableFields": ["breakdown"]}
```
