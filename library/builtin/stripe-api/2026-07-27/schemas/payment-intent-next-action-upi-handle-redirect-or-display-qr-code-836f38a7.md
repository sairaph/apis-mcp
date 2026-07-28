---
title: payment_intent_next_action_upi_handle_redirect_or_display_qr_code
page_id: schema-payment-intent-next-action-upi-handle-redirect-or-display-qr-code-836f38a7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_upi_handle_redirect_or_display_qr_code

```yaml
{"title": "PaymentIntentNextActionUpiHandleRedirectOrDisplayQrCode", "required": ["hosted_instructions_url", "qr_code"], "type": "object", "properties": {"hosted_instructions_url": {"maxLength": 5000, "type": "string", "description": "The URL to the hosted UPI instructions page, which allows customers to view the QR code."}, "qr_code": {"$ref": "#/components/schemas/payment_intent_next_action_upiqr_code"}}, "description": "", "x-expandableFields": ["qr_code"]}
```
