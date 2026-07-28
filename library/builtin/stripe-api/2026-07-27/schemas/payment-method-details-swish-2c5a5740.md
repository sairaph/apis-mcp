---
title: payment_method_details_swish
page_id: schema-payment-method-details-swish-2c5a5740
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_swish

```yaml
{"title": "payment_method_details_swish", "type": "object", "properties": {"fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies the payer's Swish account. You can use this attribute to check whether two Swish transactions were paid for by the same payer", "nullable": true}, "payment_reference": {"maxLength": 5000, "type": "string", "description": "Payer bank reference number for the payment", "nullable": true}, "verified_phone_last4": {"maxLength": 5000, "type": "string", "description": "The last four digits of the Swish account phone number", "nullable": true}}, "description": "", "x-expandableFields": []}
```
