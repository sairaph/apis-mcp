---
title: payment_flows_private_payment_methods_alipay_details
page_id: schema-payment-flows-private-payment-methods-alipay-details-32e6f416
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_private_payment_methods_alipay_details

```yaml
{"title": "PaymentFlowsPrivatePaymentMethodsAlipayDetails", "type": "object", "properties": {"buyer_id": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same."}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.", "nullable": true}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "Transaction ID of this particular Alipay transaction.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
