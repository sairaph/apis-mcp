---
title: invoice_payment_method_options_pix
page_id: schema-invoice-payment-method-options-pix-f5711833
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_pix

```yaml
{"title": "invoice_payment_method_options_pix", "type": "object", "properties": {"amount_includes_iof": {"type": "string", "description": "Determines if the amount includes the IOF tax.", "nullable": true, "enum": ["always", "never"]}, "expires_after_seconds": {"type": "integer", "description": "The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds."}}, "description": "", "x-expandableFields": []}
```
