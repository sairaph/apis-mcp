---
title: payment_links_resource_payment_intent_data
page_id: schema-payment-links-resource-payment-intent-data-4ae9ef6d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_payment_intent_data

```yaml
{"title": "PaymentLinksResourcePaymentIntentData", "required": ["metadata"], "type": "object", "properties": {"capture_method": {"type": "string", "description": "Indicates when the funds will be captured from the customer's account.", "nullable": true, "enum": ["automatic", "automatic_async", "manual"]}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link."}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with the payment method collected during checkout.", "nullable": true, "enum": ["off_session", "on_session"]}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "For a non-card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge.", "nullable": true}, "statement_descriptor_suffix": {"maxLength": 5000, "type": "string", "description": "For a card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge. Concatenated with the account's statement descriptor prefix to form the complete statement descriptor.", "nullable": true}, "transfer_group": {"maxLength": 5000, "type": "string", "description": "A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
