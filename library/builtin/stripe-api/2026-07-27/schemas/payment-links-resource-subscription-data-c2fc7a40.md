---
title: payment_links_resource_subscription_data
page_id: schema-payment-links-resource-subscription-data-c2fc7a40
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_subscription_data

```yaml
{"title": "PaymentLinksResourceSubscriptionData", "required": ["invoice_settings", "metadata"], "type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.", "nullable": true}, "invoice_settings": {"$ref": "#/components/schemas/payment_links_resource_subscription_data_invoice_settings"}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link."}, "trial_period_days": {"type": "integer", "description": "Integer representing the number of trial period days before the customer is charged for the first time.", "nullable": true}, "trial_settings": {"description": "Settings related to subscription trials.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscriptions_trials_resource_trial_settings"}]}}, "description": "", "x-expandableFields": ["invoice_settings", "trial_settings"]}
```
