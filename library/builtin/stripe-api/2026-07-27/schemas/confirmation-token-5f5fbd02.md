---
title: confirmation_token
page_id: schema-confirmation-token-5f5fbd02
path: schemas
description: |-
    ConfirmationTokens help transport client side data collected by Stripe JS over
    to your server for confirming a PaymentIntent or SetupIntent. If the confirmation
    is successful, values present on the ConfirmationToken are written onto the Intent.

    To learn more about how to use ConfirmationToken, visit the related guides:
    - [Finalize payments on the server](https://docs.stripe.com/payments/finalize-payments-on-the-server)
    - [Build two-step confirmation](https://docs.stripe.com/payments/build-a-two-step-confirmation).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_token

ConfirmationTokens help transport client side data collected by Stripe JS over
to your server for confirming a PaymentIntent or SetupIntent. If the confirmation
is successful, values present on the ConfirmationToken are written onto the Intent.

To learn more about how to use ConfirmationToken, visit the related guides:
- [Finalize payments on the server](https://docs.stripe.com/payments/finalize-payments-on-the-server)
- [Build two-step confirmation](https://docs.stripe.com/payments/build-a-two-step-confirmation).

```yaml
{"title": "ConfirmationTokensResourceConfirmationToken", "required": ["created", "id", "livemode", "object", "use_stripe_sdk"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expires_at": {"type": "integer", "description": "Time at which this ConfirmationToken expires and can no longer be used to confirm a PaymentIntent or SetupIntent.", "format": "unix-time", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "mandate_data": {"description": "Data used for generating a Mandate.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_mandate_data"}]}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["confirmation_token"]}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "ID of the PaymentIntent that this ConfirmationToken was used to confirm, or null if this ConfirmationToken has not yet been used.", "nullable": true}, "payment_method_options": {"description": "Payment-method-specific configuration for this ConfirmationToken.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_payment_method_options"}]}, "payment_method_preview": {"description": "Payment details collected by the Payment Element, used to create a PaymentMethod when a PaymentIntent or SetupIntent is confirmed with this ConfirmationToken.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_payment_method_preview"}]}, "return_url": {"maxLength": 5000, "type": "string", "description": "Return URL used to confirm the Intent.", "nullable": true}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this ConfirmationToken's payment method.\n\nThe presence of this property will [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete.", "nullable": true, "enum": ["off_session", "on_session"]}, "setup_intent": {"maxLength": 5000, "type": "string", "description": "ID of the SetupIntent that this ConfirmationToken was used to confirm, or null if this ConfirmationToken has not yet been used.", "nullable": true}, "shipping": {"description": "Shipping information collected on this ConfirmationToken.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_shipping"}]}, "use_stripe_sdk": {"type": "boolean", "description": "Indicates whether the Stripe SDK is used to handle confirmation flow. Defaults to `true` on ConfirmationToken."}}, "description": "ConfirmationTokens help transport client side data collected by Stripe JS over\nto your server for confirming a PaymentIntent or SetupIntent. If the confirmation\nis successful, values present on the ConfirmationToken are written onto the Intent.\n\nTo learn more about how to use ConfirmationToken, visit the related guides:\n- [Finalize payments on the server](https://docs.stripe.com/payments/finalize-payments-on-the-server)\n- [Build two-step confirmation](https://docs.stripe.com/payments/build-a-two-step-confirmation).", "x-expandableFields": ["mandate_data", "payment_method_options", "payment_method_preview", "shipping"], "x-resourceId": "confirmation_token"}
```
