---
title: payment_intent_next_action_pix_display_qr_code
page_id: schema-payment-intent-next-action-pix-display-qr-code-683f400b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_pix_display_qr_code

```yaml
{"title": "PaymentIntentNextActionPixDisplayQrCode", "type": "object", "properties": {"data": {"maxLength": 5000, "type": "string", "description": "The raw data string used to generate QR code, it should be used together with QR code library."}, "expires_at": {"type": "integer", "description": "The date (unix timestamp) when the PIX expires."}, "hosted_instructions_url": {"maxLength": 5000, "type": "string", "description": "The URL to the hosted pix instructions page, which allows customers to view the pix QR code."}, "image_url_png": {"maxLength": 5000, "type": "string", "description": "The image_url_png string used to render png QR code"}, "image_url_svg": {"maxLength": 5000, "type": "string", "description": "The image_url_svg string used to render svg QR code"}}, "description": "", "x-expandableFields": []}
```
