---
title: payment_method_config_resource_display_preference
page_id: schema-payment-method-config-resource-display-preference-5174a847
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_config_resource_display_preference

```yaml
{"title": "PaymentMethodConfigResourceDisplayPreference", "required": ["preference", "value"], "type": "object", "properties": {"overridable": {"type": "boolean", "description": "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.", "nullable": true}, "preference": {"type": "string", "description": "The account's display preference.", "enum": ["none", "off", "on"]}, "value": {"type": "string", "description": "The effective display preference value.", "enum": ["off", "on"]}}, "description": "", "x-expandableFields": []}
```
