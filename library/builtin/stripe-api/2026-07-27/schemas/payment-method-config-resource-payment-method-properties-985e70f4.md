---
title: payment_method_config_resource_payment_method_properties
page_id: schema-payment-method-config-resource-payment-method-properties-985e70f4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_config_resource_payment_method_properties

```yaml
{"title": "PaymentMethodConfigResourcePaymentMethodProperties", "required": ["available", "display_preference"], "type": "object", "properties": {"available": {"type": "boolean", "description": "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active."}, "display_preference": {"$ref": "#/components/schemas/payment_method_config_resource_display_preference"}}, "description": "", "x-expandableFields": ["display_preference"]}
```
