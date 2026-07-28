---
title: payment_intent_next_action_swish_qr_code
page_id: schema-payment-intent-next-action-swish-qr-code-c9d38186
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_swish_qr_code

```yaml
{"title": "PaymentIntentNextActionSwishQRCode", "required": ["data", "image_url_png", "image_url_svg"], "type": "object", "properties": {"data": {"maxLength": 5000, "type": "string", "description": "The raw data string used to generate QR code, it should be used together with QR code library."}, "image_url_png": {"maxLength": 5000, "type": "string", "description": "The image_url_png string used to render QR code"}, "image_url_svg": {"maxLength": 5000, "type": "string", "description": "The image_url_svg string used to render QR code"}}, "description": "", "x-expandableFields": []}
```
