---
title: payment_intent_next_action_paynow_display_qr_code
page_id: schema-payment-intent-next-action-paynow-display-qr-code-0de1c752
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_paynow_display_qr_code

```yaml
{"title": "PaymentIntentNextActionPaynowDisplayQrCode", "required": ["data", "image_url_png", "image_url_svg"], "type": "object", "properties": {"data": {"maxLength": 5000, "type": "string", "description": "The raw data string used to generate QR code, it should be used together with QR code library."}, "hosted_instructions_url": {"maxLength": 5000, "type": "string", "description": "The URL to the hosted PayNow instructions page, which allows customers to view the PayNow QR code.", "nullable": true}, "image_url_png": {"maxLength": 5000, "type": "string", "description": "The image_url_png string used to render QR code"}, "image_url_svg": {"maxLength": 5000, "type": "string", "description": "The image_url_svg string used to render QR code"}}, "description": "", "x-expandableFields": []}
```
