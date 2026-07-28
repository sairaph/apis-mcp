---
title: payment_method_details_payment_record_sepa_debit
page_id: schema-payment-method-details-payment-record-sepa-debit-6c7974f4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_sepa_debit

```yaml
{"title": "payment_method_details_payment_record_sepa_debit", "type": "object", "properties": {"bank_code": {"maxLength": 5000, "type": "string", "description": "Bank code of bank associated with the bank account.", "nullable": true}, "branch_code": {"maxLength": 5000, "type": "string", "description": "Branch code of bank associated with the bank account.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country the bank account is located in.", "nullable": true}, "expected_debit_date": {"maxLength": 5000, "type": "string", "description": "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format."}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four characters of the IBAN.", "nullable": true}, "mandate": {"maxLength": 5000, "type": "string", "description": "Find the ID of the mandate used for this payment under the [payment_method_details.sepa_debit.mandate](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-sepa_debit-mandate) property on the Charge. Use this mandate ID to [retrieve the Mandate](https://docs.stripe.com/api/mandates/retrieve).", "nullable": true}}, "description": "", "x-expandableFields": []}
```
