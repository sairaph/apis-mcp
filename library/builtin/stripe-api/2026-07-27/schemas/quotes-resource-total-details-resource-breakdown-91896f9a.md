---
title: quotes_resource_total_details_resource_breakdown
page_id: schema-quotes-resource-total-details-resource-breakdown-91896f9a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_total_details_resource_breakdown

```yaml
{"title": "QuotesResourceTotalDetailsResourceBreakdown", "required": ["discounts", "taxes"], "type": "object", "properties": {"discounts": {"type": "array", "description": "The aggregated discounts.", "items": {"$ref": "#/components/schemas/line_items_discount_amount"}}, "taxes": {"type": "array", "description": "The aggregated tax amounts by rate.", "items": {"$ref": "#/components/schemas/line_items_tax_amount"}}}, "description": "", "x-expandableFields": ["discounts", "taxes"]}
```
