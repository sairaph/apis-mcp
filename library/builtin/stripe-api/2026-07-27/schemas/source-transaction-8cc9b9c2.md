---
title: source_transaction
page_id: schema-source-transaction-8cc9b9c2
path: schemas
description: |-
    Some payment methods have no required amount that a customer must send.
    Customers can be instructed to send any amount, and it can be made up of
    multiple transactions. As such, sources can have multiple associated
    transactions.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_transaction

Some payment methods have no required amount that a customer must send.
Customers can be instructed to send any amount, and it can be made up of
multiple transactions. As such, sources can have multiple associated
transactions.

```yaml
{"title": "SourceTransaction", "required": ["amount", "created", "currency", "id", "livemode", "object", "source", "status", "type"], "type": "object", "properties": {"ach_credit_transfer": {"$ref": "#/components/schemas/source_transaction_ach_credit_transfer_data"}, "amount": {"type": "integer", "description": "A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount your customer has pushed to the receiver."}, "chf_credit_transfer": {"$ref": "#/components/schemas/source_transaction_chf_credit_transfer_data"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "gbp_credit_transfer": {"$ref": "#/components/schemas/source_transaction_gbp_credit_transfer_data"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["source_transaction"]}, "paper_check": {"$ref": "#/components/schemas/source_transaction_paper_check_data"}, "sepa_credit_transfer": {"$ref": "#/components/schemas/source_transaction_sepa_credit_transfer_data"}, "source": {"maxLength": 5000, "type": "string", "description": "The ID of the source this transaction is attached to."}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the transaction, one of `succeeded`, `pending`, or `failed`."}, "type": {"type": "string", "description": "The type of source this transaction is attached to.", "enum": ["ach_credit_transfer", "ach_debit", "alipay", "bancontact", "card", "card_present", "eps", "giropay", "ideal", "klarna", "multibanco", "p24", "sepa_debit", "sofort", "three_d_secure", "wechat"]}}, "description": "Some payment methods have no required amount that a customer must send.\nCustomers can be instructed to send any amount, and it can be made up of\nmultiple transactions. As such, sources can have multiple associated\ntransactions.", "x-expandableFields": ["ach_credit_transfer", "chf_credit_transfer", "gbp_credit_transfer", "paper_check", "sepa_credit_transfer"], "x-resourceId": "source_transaction"}
```
