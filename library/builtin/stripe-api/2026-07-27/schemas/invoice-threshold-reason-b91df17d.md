---
title: invoice_threshold_reason
page_id: schema-invoice-threshold-reason-b91df17d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_threshold_reason

```yaml
{"title": "InvoiceThresholdReason", "required": ["item_reasons"], "type": "object", "properties": {"amount_gte": {"type": "integer", "description": "The total invoice amount threshold boundary if it triggered the threshold invoice.", "nullable": true}, "item_reasons": {"type": "array", "description": "Indicates which line items triggered a threshold invoice.", "items": {"$ref": "#/components/schemas/invoice_item_threshold_reason"}}}, "description": "", "x-expandableFields": ["item_reasons"]}
```
