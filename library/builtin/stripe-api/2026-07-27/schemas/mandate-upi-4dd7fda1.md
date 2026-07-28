---
title: mandate_upi
page_id: schema-mandate-upi-4dd7fda1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_upi

```yaml
{"title": "mandate_upi", "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be charged for future payments.", "nullable": true}, "amount_type": {"type": "string", "description": "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.", "nullable": true, "enum": ["fixed", "maximum"]}, "description": {"maxLength": 20, "type": "string", "description": "A description of the mandate or subscription that is meant to be displayed to the customer.", "nullable": true}, "end_date": {"type": "integer", "description": "End date of the mandate or subscription.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
