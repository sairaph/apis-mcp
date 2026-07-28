---
title: quotes_resource_subscription_data_subscription_data
page_id: schema-quotes-resource-subscription-data-subscription-data-6e3e5c89
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_subscription_data_subscription_data

```yaml
{"title": "QuotesResourceSubscriptionDataSubscriptionData", "required": ["billing_mode"], "type": "object", "properties": {"billing_mode": {"$ref": "#/components/schemas/quotes_resource_subscription_data_billing_mode"}, "description": {"maxLength": 5000, "type": "string", "description": "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.", "nullable": true}, "effective_date": {"type": "integer", "description": "When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. This date is ignored if it is in the past when the quote is accepted. Measured in seconds since the Unix epoch.", "format": "unix-time", "nullable": true}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on the subscription or subscription schedule when the quote is accepted. If a recurring price is included in `line_items`, this field will be passed to the resulting subscription's `metadata` field. If `subscription_data.effective_date` is used, this field will be passed to the resulting subscription schedule's `phases.metadata` field. Unlike object-level metadata, this field is declarative. Updates will clear prior values.", "nullable": true}, "trial_period_days": {"type": "integer", "description": "Integer representing the number of trial period days before the customer is charged for the first time.", "nullable": true}}, "description": "", "x-expandableFields": ["billing_mode"]}
```
