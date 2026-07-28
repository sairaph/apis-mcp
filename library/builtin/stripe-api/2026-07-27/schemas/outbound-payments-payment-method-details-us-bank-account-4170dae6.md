---
title: outbound_payments_payment_method_details_us_bank_account
page_id: schema-outbound-payments-payment-method-details-us-bank-account-4170dae6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# outbound_payments_payment_method_details_us_bank_account

```yaml
{"title": "outbound_payments_payment_method_details_us_bank_account", "required": ["network"], "type": "object", "properties": {"account_holder_type": {"type": "string", "description": "Account holder type: individual or company.", "nullable": true, "enum": ["company", "individual"]}, "account_type": {"type": "string", "description": "Account type: checkings or savings. Defaults to checking if omitted.", "nullable": true, "enum": ["checking", "savings"]}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "mandate": {"description": "ID of the mandate used to make this payment.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/mandate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/mandate"}]}}, "network": {"type": "string", "description": "The network rails used. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.", "enum": ["ach", "us_domestic_wire"]}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing number of the bank account.", "nullable": true}}, "description": "", "x-expandableFields": ["mandate"]}
```
