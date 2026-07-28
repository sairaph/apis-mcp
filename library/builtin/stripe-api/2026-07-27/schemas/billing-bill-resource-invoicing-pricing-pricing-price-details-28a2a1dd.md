---
title: billing_bill_resource_invoicing_pricing_pricing_price_details
page_id: schema-billing-bill-resource-invoicing-pricing-pricing-price-details-28a2a1dd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_pricing_pricing_price_details

```yaml
{"title": "BillingBillResourceInvoicingPricingPricingPriceDetails", "required": ["price", "product"], "type": "object", "properties": {"price": {"description": "The ID of the price this item is associated with.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/price"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/price"}]}}, "product": {"maxLength": 5000, "type": "string", "description": "The ID of the product this item is associated with."}}, "description": "", "x-expandableFields": ["price"]}
```
