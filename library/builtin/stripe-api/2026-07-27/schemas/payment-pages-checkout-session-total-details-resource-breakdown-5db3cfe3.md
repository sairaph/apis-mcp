---
title: payment_pages_checkout_session_total_details_resource_breakdown
page_id: schema-payment-pages-checkout-session-total-details-resource-breakdown-5db3cfe3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_total_details_resource_breakdown

```yaml
{"title": "PaymentPagesCheckoutSessionTotalDetailsResourceBreakdown", "required": ["discounts", "taxes"], "type": "object", "properties": {"discounts": {"type": "array", "description": "The aggregated discounts.", "items": {"$ref": "#/components/schemas/line_items_discount_amount"}}, "taxes": {"type": "array", "description": "The aggregated tax amounts by rate.", "items": {"$ref": "#/components/schemas/line_items_tax_amount"}}}, "description": "", "x-expandableFields": ["discounts", "taxes"]}
```
