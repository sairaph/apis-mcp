---
title: payment_method_us_bank_account
page_id: schema-payment-method-us-bank-account-de1a6d70
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_us_bank_account

```yaml
{"title": "payment_method_us_bank_account", "type": "object", "properties": {"account_holder_type": {"type": "string", "description": "Account holder type: individual or company.", "nullable": true, "enum": ["company", "individual"]}, "account_type": {"type": "string", "description": "Account type: checkings or savings. Defaults to checking if omitted.", "nullable": true, "enum": ["checking", "savings"]}, "bank_name": {"maxLength": 5000, "type": "string", "description": "The name of the bank.", "nullable": true}, "financial_connections_account": {"maxLength": 5000, "type": "string", "description": "The ID of the Financial Connections Account used to create the payment method.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "networks": {"description": "Contains information about US bank account networks that can be used.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/us_bank_account_networks"}]}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing number of the bank account.", "nullable": true}, "status_details": {"description": "Contains information about the future reusability of this PaymentMethod.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_method_us_bank_account_status_details"}]}}, "description": "", "x-expandableFields": ["networks", "status_details"]}
```
