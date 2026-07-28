---
title: payment_intent_next_action_redirect_to_url
page_id: schema-payment-intent-next-action-redirect-to-url-4a12b7e0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_redirect_to_url

```yaml
{"title": "PaymentIntentNextActionRedirectToUrl", "type": "object", "properties": {"return_url": {"maxLength": 5000, "type": "string", "description": "If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "The URL you must redirect your customer to in order to authenticate the payment.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
