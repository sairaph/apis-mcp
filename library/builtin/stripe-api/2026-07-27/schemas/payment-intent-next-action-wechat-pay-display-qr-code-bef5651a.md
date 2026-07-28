---
title: payment_intent_next_action_wechat_pay_display_qr_code
page_id: schema-payment-intent-next-action-wechat-pay-display-qr-code-bef5651a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_wechat_pay_display_qr_code

```yaml
{"title": "PaymentIntentNextActionWechatPayDisplayQrCode", "required": ["data", "hosted_instructions_url", "image_data_url", "image_url_png", "image_url_svg"], "type": "object", "properties": {"data": {"maxLength": 5000, "type": "string", "description": "The data being used to generate QR code"}, "hosted_instructions_url": {"maxLength": 5000, "type": "string", "description": "The URL to the hosted WeChat Pay instructions page, which allows customers to view the WeChat Pay QR code."}, "image_data_url": {"maxLength": 5000, "type": "string", "description": "The base64 image data for a pre-generated QR code"}, "image_url_png": {"maxLength": 5000, "type": "string", "description": "The image_url_png string used to render QR code"}, "image_url_svg": {"maxLength": 5000, "type": "string", "description": "The image_url_svg string used to render QR code"}}, "description": "", "x-expandableFields": []}
```
