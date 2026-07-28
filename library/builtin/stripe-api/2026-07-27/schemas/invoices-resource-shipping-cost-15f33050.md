---
title: invoices_resource_shipping_cost
page_id: schema-invoices-resource-shipping-cost-15f33050
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_shipping_cost

```yaml
{"title": "InvoicesResourceShippingCost", "required": ["amount_subtotal", "amount_tax", "amount_total"], "type": "object", "properties": {"amount_subtotal": {"type": "integer", "description": "Total shipping cost before any taxes are applied."}, "amount_tax": {"type": "integer", "description": "Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0."}, "amount_total": {"type": "integer", "description": "Total shipping cost after taxes are applied."}, "shipping_rate": {"description": "The ID of the ShippingRate for this invoice.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/shipping_rate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/shipping_rate"}]}}, "taxes": {"type": "array", "description": "The taxes applied to the shipping rate.", "items": {"$ref": "#/components/schemas/line_items_tax_amount"}}}, "description": "", "x-expandableFields": ["shipping_rate", "taxes"]}
```
