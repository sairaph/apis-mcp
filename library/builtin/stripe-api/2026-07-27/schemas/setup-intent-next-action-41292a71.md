---
title: setup_intent_next_action
page_id: schema-setup-intent-next-action-41292a71
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_intent_next_action

```yaml
{"title": "SetupIntentNextAction", "required": ["type"], "type": "object", "properties": {"blik_authorize": {"$ref": "#/components/schemas/payment_intent_next_action_blik_authorize"}, "cashapp_handle_redirect_or_display_qr_code": {"$ref": "#/components/schemas/payment_intent_next_action_cashapp_handle_redirect_or_display_qr_code"}, "pix_display_qr_code": {"$ref": "#/components/schemas/setup_intent_next_action_pix_display_qr_code"}, "redirect_to_url": {"$ref": "#/components/schemas/setup_intent_next_action_redirect_to_url"}, "type": {"maxLength": 5000, "type": "string", "description": "Type of the next action to perform. Refer to the other child attributes under `next_action` for available values. Examples include: `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`."}, "upi_handle_redirect_or_display_qr_code": {"$ref": "#/components/schemas/payment_intent_next_action_upi_handle_redirect_or_display_qr_code"}, "use_stripe_sdk": {"type": "object", "description": "When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js."}, "verify_with_microdeposits": {"$ref": "#/components/schemas/setup_intent_next_action_verify_with_microdeposits"}}, "description": "", "x-expandableFields": ["blik_authorize", "cashapp_handle_redirect_or_display_qr_code", "pix_display_qr_code", "redirect_to_url", "upi_handle_redirect_or_display_qr_code", "verify_with_microdeposits"]}
```
