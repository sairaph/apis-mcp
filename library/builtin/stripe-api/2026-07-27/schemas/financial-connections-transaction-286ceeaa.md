---
title: financial_connections.transaction
page_id: schema-financial-connections-transaction-286ceeaa
path: schemas
description: A Transaction represents a real transaction that affects a Financial Connections Account balance.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# financial_connections.transaction

A Transaction represents a real transaction that affects a Financial Connections Account balance.

```yaml
{"title": "BankConnectionsResourceTransaction", "required": ["account", "amount", "currency", "description", "id", "livemode", "object", "status", "status_transitions", "transacted_at", "transaction_refresh", "updated"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string", "description": "The ID of the Financial Connections Account this transaction belongs to."}, "amount": {"type": "integer", "description": "The amount of this transaction, in cents (or local equivalent)."}, "currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "description": {"maxLength": 5000, "type": "string", "description": "The description of this transaction."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["financial_connections.transaction"]}, "status": {"type": "string", "description": "The status of the transaction.", "enum": ["pending", "posted", "void"]}, "status_transitions": {"$ref": "#/components/schemas/bank_connections_resource_transaction_resource_status_transitions"}, "transacted_at": {"type": "integer", "description": "Time at which the transaction was transacted. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "transaction_refresh": {"maxLength": 5000, "type": "string", "description": "The token of the transaction refresh that last updated or created this transaction."}, "updated": {"type": "integer", "description": "Time at which the object was last updated. Measured in seconds since the Unix epoch.", "format": "unix-time"}}, "description": "A Transaction represents a real transaction that affects a Financial Connections Account balance.", "x-expandableFields": ["status_transitions"], "x-resourceId": "financial_connections.transaction"}
```
