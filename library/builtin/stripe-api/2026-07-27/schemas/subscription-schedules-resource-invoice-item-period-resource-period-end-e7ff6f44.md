---
title: subscription_schedules_resource_invoice_item_period_resource_period_end
page_id: schema-subscription-schedules-resource-invoice-item-period-resource-period-end-e7ff6f44
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedules_resource_invoice_item_period_resource_period_end

```yaml
{"title": "SubscriptionSchedulesResourceInvoiceItemPeriodResourcePeriodEnd", "required": ["type"], "type": "object", "properties": {"timestamp": {"type": "integer", "description": "A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.", "format": "unix-time"}, "type": {"type": "string", "description": "Select how to calculate the end of the invoice item period.", "enum": ["min_item_period_end", "phase_end", "timestamp"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
