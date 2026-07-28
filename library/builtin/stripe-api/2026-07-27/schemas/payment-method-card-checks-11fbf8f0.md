---
title: payment_method_card_checks
page_id: schema-payment-method-card-checks-11fbf8f0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_card_checks

```yaml
{"title": "payment_method_card_checks", "type": "object", "properties": {"address_line1_check": {"maxLength": 5000, "type": "string", "description": "If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true}, "address_postal_code_check": {"maxLength": 5000, "type": "string", "description": "If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true}, "cvc_check": {"maxLength": 5000, "type": "string", "description": "If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
