---
title: subscription_pending_invoice_item_interval
page_id: schema-subscription-pending-invoice-item-interval-642c6850
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_pending_invoice_item_interval

```yaml
{"title": "SubscriptionPendingInvoiceItemInterval", "required": ["interval", "interval_count"], "type": "object", "properties": {"interval": {"type": "string", "description": "Specifies invoicing frequency. Either `day`, `week`, `month` or `year`.", "enum": ["day", "month", "week", "year"]}, "interval_count": {"type": "integer", "description": "The number of intervals between invoices. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks)."}}, "description": "", "x-expandableFields": []}
```
