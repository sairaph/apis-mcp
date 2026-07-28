---
title: invoice_mandate_options_card
page_id: schema-invoice-mandate-options-card-b6baa8d8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_mandate_options_card

```yaml
{"title": "invoice_mandate_options_card", "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be charged for future payments, specified in the presentment currency.", "nullable": true}, "amount_type": {"type": "string", "description": "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.", "nullable": true, "enum": ["fixed", "maximum"]}, "description": {"maxLength": 200, "type": "string", "description": "A description of the mandate or subscription that is meant to be displayed to the customer.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
