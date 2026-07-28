---
title: payment_intent_next_action_wechat_pay_redirect_to_android_app
page_id: schema-payment-intent-next-action-wechat-pay-redirect-to-android-app-a7221ee7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_wechat_pay_redirect_to_android_app

```yaml
{"title": "PaymentIntentNextActionWechatPayRedirectToAndroidApp", "required": ["app_id", "nonce_str", "package", "partner_id", "prepay_id", "sign", "timestamp"], "type": "object", "properties": {"app_id": {"maxLength": 5000, "type": "string", "description": "app_id is the APP ID registered on WeChat open platform"}, "nonce_str": {"maxLength": 5000, "type": "string", "description": "nonce_str is a random string"}, "package": {"maxLength": 5000, "type": "string", "description": "package is static value"}, "partner_id": {"maxLength": 5000, "type": "string", "description": "an unique merchant ID assigned by WeChat Pay"}, "prepay_id": {"maxLength": 5000, "type": "string", "description": "an unique trading ID assigned by WeChat Pay"}, "sign": {"maxLength": 5000, "type": "string", "description": "A signature"}, "timestamp": {"maxLength": 5000, "type": "string", "description": "Specifies the current time in epoch format"}}, "description": "", "x-expandableFields": []}
```
