---
title: payment_links_resource_shipping_option
page_id: schema-payment-links-resource-shipping-option-c4b1fd94
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_shipping_option

```yaml
{"title": "PaymentLinksResourceShippingOption", "required": ["shipping_amount", "shipping_rate"], "type": "object", "properties": {"shipping_amount": {"type": "integer", "description": "A non-negative integer in cents representing how much to charge."}, "shipping_rate": {"description": "The ID of the Shipping Rate to use for this shipping option.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/shipping_rate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/shipping_rate"}]}}}, "description": "", "x-expandableFields": ["shipping_rate"]}
```
