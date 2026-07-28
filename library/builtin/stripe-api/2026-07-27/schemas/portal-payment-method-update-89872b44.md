---
title: portal_payment_method_update
page_id: schema-portal-payment-method-update-89872b44
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_payment_method_update

```yaml
{"title": "PortalPaymentMethodUpdate", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the feature is enabled."}, "payment_method_configuration": {"maxLength": 5000, "type": "string", "description": "The [Payment Method Configuration](/api/payment_method_configurations) to use for this portal session. When specified, customers will be able to update their payment method to one of the options specified by the payment method configuration. If not set, the default payment method configuration is used.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
