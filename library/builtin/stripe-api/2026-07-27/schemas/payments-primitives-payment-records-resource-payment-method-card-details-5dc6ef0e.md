---
title: payments_primitives_payment_records_resource_payment_method_card_details_resource_checks
page_id: schema-payments-primitives-payment-records-resource-payment-method-card-details-5dc6ef0e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_card_details_resource_checks

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCardDetailsResourceChecks", "type": "object", "properties": {"address_line1_check": {"type": "string", "description": "If you provide a value for `address.line1`, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true, "enum": ["fail", "pass", "unavailable", "unchecked"]}, "address_postal_code_check": {"type": "string", "description": "If you provide a address postal code, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true, "enum": ["fail", "pass", "unavailable", "unchecked"]}, "cvc_check": {"type": "string", "description": "If you provide a CVC, the check results is one of `pass`, `fail`, `unavailable`, or `unchecked`.", "nullable": true, "enum": ["fail", "pass", "unavailable", "unchecked"]}}, "description": "", "x-expandableFields": []}
```
