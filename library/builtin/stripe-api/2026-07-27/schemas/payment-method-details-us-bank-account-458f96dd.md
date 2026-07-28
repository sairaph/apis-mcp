---
title: payment_method_details_us_bank_account
page_id: schema-payment-method-details-us-bank-account-458f96dd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_us_bank_account

```yaml
{"title": "payment_method_details_us_bank_account", "type": "object", "properties": {"account_holder_type": {"type": "string", "description": "Account holder type: individual or company.", "nullable": true, "enum": ["company", "individual"]}, "account_type": {"type": "string", "description": "Account type: checkings or savings. Defaults to checking if omitted.", "nullable": true, "enum": ["checking", "savings"]}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "expected_debit_date": {"maxLength": 5000, "type": "string", "description": "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format."}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "mandate": {"description": "ID of the mandate used to make this payment.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/mandate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/mandate"}]}}, "payment_reference": {"maxLength": 5000, "type": "string", "description": "Reference number to locate ACH payments with customer's bank.", "nullable": true}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing number of the bank account.", "nullable": true}}, "description": "", "x-expandableFields": ["mandate"]}
```
