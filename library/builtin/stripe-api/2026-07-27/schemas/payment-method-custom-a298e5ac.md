---
title: payment_method_custom
page_id: schema-payment-method-custom-a298e5ac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_custom

```yaml
{"title": "payment_method_custom", "required": ["type"], "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string", "description": "Display name of the Dashboard-only CustomPaymentMethodType.", "nullable": true}, "logo": {"description": "Contains information about the Dashboard-only CustomPaymentMethodType logo.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/custom_logo"}]}, "type": {"maxLength": 5000, "type": "string", "description": "ID of the Dashboard-only CustomPaymentMethodType. Not expandable."}}, "description": "", "x-expandableFields": ["logo"]}
```
