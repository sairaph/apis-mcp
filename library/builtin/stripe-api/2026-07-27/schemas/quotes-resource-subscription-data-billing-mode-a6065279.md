---
title: quotes_resource_subscription_data_billing_mode
page_id: schema-quotes-resource-subscription-data-billing-mode-a6065279
path: schemas
description: The billing mode of the quote.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_subscription_data_billing_mode

The billing mode of the quote.

```yaml
{"title": "QuotesResourceSubscriptionDataBillingMode", "required": ["type"], "type": "object", "properties": {"flexible": {"$ref": "#/components/schemas/subscriptions_resource_billing_mode_flexible"}, "type": {"type": "string", "description": "Controls how prorations and invoices for subscriptions are calculated and orchestrated.", "enum": ["classic", "flexible"]}}, "description": "The billing mode of the quote.", "x-expandableFields": ["flexible"]}
```
