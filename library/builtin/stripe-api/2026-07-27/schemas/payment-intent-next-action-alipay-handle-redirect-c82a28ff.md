---
title: payment_intent_next_action_alipay_handle_redirect
page_id: schema-payment-intent-next-action-alipay-handle-redirect-c82a28ff
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_alipay_handle_redirect

```yaml
{"title": "PaymentIntentNextActionAlipayHandleRedirect", "type": "object", "properties": {"native_data": {"maxLength": 5000, "type": "string", "description": "The native data to be used with Alipay SDK you must redirect your customer to in order to authenticate the payment in an Android App.", "nullable": true}, "native_url": {"maxLength": 5000, "type": "string", "description": "The native URL you must redirect your customer to in order to authenticate the payment in an iOS App.", "nullable": true}, "return_url": {"maxLength": 5000, "type": "string", "description": "If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "The URL you must redirect your customer to in order to authenticate the payment.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
