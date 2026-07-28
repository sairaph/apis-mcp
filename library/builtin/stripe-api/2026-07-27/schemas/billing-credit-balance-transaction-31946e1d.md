---
title: billing.credit_balance_transaction
page_id: schema-billing-credit-balance-transaction-31946e1d
path: schemas
description: A credit balance transaction is a resource representing a transaction (either a credit or a debit) against an existing credit grant.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.credit_balance_transaction

A credit balance transaction is a resource representing a transaction (either a credit or a debit) against an existing credit grant.

```yaml
{"title": "CreditBalanceTransaction", "required": ["created", "credit_grant", "effective_at", "id", "livemode", "object"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "credit": {"description": "Credit details for this credit balance transaction. Only present if type is `credit`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_credit_grants_resource_balance_credit"}]}, "credit_grant": {"description": "The credit grant associated with this credit balance transaction.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/billing.credit_grant"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/billing.credit_grant"}]}}, "debit": {"description": "Debit details for this credit balance transaction. Only present if type is `debit`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_credit_grants_resource_balance_debit"}]}, "effective_at": {"type": "integer", "description": "The effective time of this credit balance transaction.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.credit_balance_transaction"]}, "test_clock": {"description": "ID of the test clock this credit balance transaction belongs to.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/test_helpers.test_clock"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/test_helpers.test_clock"}]}}, "type": {"type": "string", "description": "The type of credit balance transaction (credit or debit).", "nullable": true, "enum": ["credit", "debit"]}}, "description": "A credit balance transaction is a resource representing a transaction (either a credit or a debit) against an existing credit grant.", "x-expandableFields": ["credit", "credit_grant", "debit", "test_clock"], "x-resourceId": "billing.credit_balance_transaction"}
```
