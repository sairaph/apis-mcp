---
title: quotes_resource_recurring
page_id: schema-quotes-resource-recurring-1e92fed1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_recurring

```yaml
{"title": "QuotesResourceRecurring", "required": ["amount_subtotal", "amount_total", "interval", "interval_count", "total_details"], "type": "object", "properties": {"amount_subtotal": {"type": "integer", "description": "Total before any discounts or taxes are applied."}, "amount_total": {"type": "integer", "description": "Total after discounts and taxes are applied."}, "interval": {"type": "string", "description": "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.", "enum": ["day", "month", "week", "year"]}, "interval_count": {"type": "integer", "description": "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months."}, "total_details": {"$ref": "#/components/schemas/quotes_resource_total_details"}}, "description": "", "x-expandableFields": ["total_details"]}
```
