---
title: payment_method_details_payment_record_acss_debit
page_id: schema-payment-method-details-payment-record-acss-debit-9d57e01e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_acss_debit

```yaml
{"title": "payment_method_details_payment_record_acss_debit", "type": "object", "properties": {"bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "expected_debit_date": {"maxLength": 5000, "type": "string", "description": "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format."}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "institution_number": {"maxLength": 5000, "type": "string", "description": "Institution number of the bank account", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "mandate": {"maxLength": 5000, "type": "string", "description": "ID of the mandate used to make this payment."}, "transit_number": {"maxLength": 5000, "type": "string", "description": "Transit number of the bank account.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
