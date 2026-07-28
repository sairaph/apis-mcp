---
title: bank_connections_resource_balance
page_id: schema-bank-connections-resource-balance-3d874682
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_balance

```yaml
{"title": "BankConnectionsResourceBalance", "required": ["as_of", "current", "type"], "type": "object", "properties": {"as_of": {"type": "integer", "description": "The time that the external institution calculated this balance. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "cash": {"$ref": "#/components/schemas/bank_connections_resource_balance_api_resource_cash_balance"}, "credit": {"$ref": "#/components/schemas/bank_connections_resource_balance_api_resource_credit_balance"}, "current": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "The balances owed to (or by) the account holder, before subtracting any outbound pending transactions or adding any inbound pending transactions.\n\nEach key is a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.\n\nEach value is a integer amount. A positive amount indicates money owed to the account holder. A negative amount indicates money owed by the account holder."}, "type": {"type": "string", "description": "The `type` of the balance. An additional hash is included on the balance with a name matching this value.", "enum": ["cash", "credit"]}}, "description": "", "x-expandableFields": ["cash", "credit"]}
```
