---
title: dispute_payment_method_details
page_id: schema-dispute-payment-method-details-6a3d1098
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_payment_method_details

```yaml
{"title": "DisputePaymentMethodDetails", "required": ["type"], "type": "object", "properties": {"amazon_pay": {"$ref": "#/components/schemas/dispute_payment_method_details_amazon_pay"}, "card": {"$ref": "#/components/schemas/dispute_payment_method_details_card"}, "klarna": {"$ref": "#/components/schemas/dispute_payment_method_details_klarna"}, "paypal": {"$ref": "#/components/schemas/dispute_payment_method_details_paypal"}, "type": {"type": "string", "description": "Payment method type.", "enum": ["amazon_pay", "card", "klarna", "paypal"]}}, "description": "", "x-expandableFields": ["amazon_pay", "card", "klarna", "paypal"]}
```
