---
title: payment_intent_next_action_klarna_display_qr_code
page_id: schema-payment-intent-next-action-klarna-display-qr-code-f8b6e782
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_klarna_display_qr_code

```yaml
{"title": "PaymentIntentNextActionKlarnaDisplayQrCode", "required": ["data", "image_url_png", "image_url_svg"], "type": "object", "properties": {"data": {"maxLength": 5000, "type": "string", "description": "The data being used to generate QR code"}, "expires_at": {"type": "integer", "description": "The timestamp at which the QR code expires.", "format": "unix-time", "nullable": true}, "image_url_png": {"maxLength": 5000, "type": "string", "description": "The image_url_png string used to render QR code"}, "image_url_svg": {"maxLength": 5000, "type": "string", "description": "The image_url_svg string used to render QR code"}}, "description": "", "x-expandableFields": []}
```
