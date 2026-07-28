---
title: invoice_line_item_period
page_id: schema-invoice-line-item-period-ee91e8dc
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_line_item_period

```yaml
{"title": "InvoiceLineItemPeriod", "required": ["end", "start"], "type": "object", "properties": {"end": {"type": "integer", "description": "The end of the period, which must be greater than or equal to the start. This value is inclusive.", "format": "unix-time"}, "start": {"type": "integer", "description": "The start of the period. This value is inclusive.", "format": "unix-time"}}, "description": "", "x-expandableFields": []}
```
