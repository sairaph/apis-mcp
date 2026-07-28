---
title: payment_pages_checkout_session_shipping_cost
page_id: schema-payment-pages-checkout-session-shipping-cost-1e27d7ab
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_shipping_cost

```yaml
{"title": "PaymentPagesCheckoutSessionShippingCost", "required": ["amount_subtotal", "amount_tax", "amount_total"], "type": "object", "properties": {"amount_subtotal": {"type": "integer", "description": "Total shipping cost before any discounts or taxes are applied."}, "amount_tax": {"type": "integer", "description": "Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0."}, "amount_total": {"type": "integer", "description": "Total shipping cost after discounts and taxes are applied."}, "shipping_rate": {"description": "The ID of the ShippingRate for this order.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/shipping_rate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/shipping_rate"}]}}, "taxes": {"type": "array", "description": "The taxes applied to the shipping rate.", "items": {"$ref": "#/components/schemas/line_items_tax_amount"}}}, "description": "", "x-expandableFields": ["shipping_rate", "taxes"]}
```
