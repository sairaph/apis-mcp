---
title: subscription_schedules_resource_invoice_item_period_resource_period_start
page_id: schema-subscription-schedules-resource-invoice-item-period-resource-period-star-304c2e75
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedules_resource_invoice_item_period_resource_period_start

```yaml
{"title": "SubscriptionSchedulesResourceInvoiceItemPeriodResourcePeriodStart", "required": ["type"], "type": "object", "properties": {"timestamp": {"type": "integer", "description": "A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.", "format": "unix-time"}, "type": {"type": "string", "description": "Select how to calculate the start of the invoice item period.", "enum": ["max_item_period_start", "phase_start", "timestamp"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
