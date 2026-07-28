---
title: proration_details
page_id: schema-proration-details-637afd54
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# proration_details

```yaml
{"title": "ProrationDetails", "required": ["discount_amounts"], "type": "object", "properties": {"credited_items": {"description": "For a credit proration, links to the debit invoice line items or invoice item that the credit applies to.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/invoice_item_proration_credited_items"}]}, "discount_amounts": {"type": "array", "description": "Discount amounts applied when the proration was created.", "items": {"$ref": "#/components/schemas/discounts_resource_discount_amount"}}}, "description": "", "x-expandableFields": ["credited_items", "discount_amounts"]}
```
