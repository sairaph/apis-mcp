---
title: subscription_payment_method_options_pix
page_id: schema-subscription-payment-method-options-pix-1bbbe06a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_payment_method_options_pix

```yaml
{"title": "subscription_payment_method_options_pix", "type": "object", "properties": {"expires_after_seconds": {"type": "integer", "description": "The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds."}, "mandate_options": {"$ref": "#/components/schemas/subscription_payment_method_options_mandate_options_pix"}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
