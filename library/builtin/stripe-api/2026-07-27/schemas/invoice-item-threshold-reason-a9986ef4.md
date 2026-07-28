---
title: invoice_item_threshold_reason
page_id: schema-invoice-item-threshold-reason-a9986ef4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_item_threshold_reason

```yaml
{"title": "InvoiceItemThresholdReason", "required": ["line_item_ids", "usage_gte"], "type": "object", "properties": {"line_item_ids": {"type": "array", "description": "The IDs of the line items that triggered the threshold invoice.", "items": {"maxLength": 5000, "type": "string"}}, "usage_gte": {"type": "integer", "description": "The quantity threshold boundary that applied to the given line item."}}, "description": "", "x-expandableFields": []}
```
