---
title: bank_connections_resource_balance_api_resource_cash_balance
page_id: schema-bank-connections-resource-balance-api-resource-cash-balance-ec9c15de
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_balance_api_resource_cash_balance

```yaml
{"title": "BankConnectionsResourceBalanceAPIResourceCashBalance", "type": "object", "properties": {"available": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "The funds available to the account holder. Typically this is the current balance after subtracting any outbound pending transactions and adding any inbound pending transactions.\n\nEach key is a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.\n\nEach value is a integer amount. A positive amount indicates money owed to the account holder. A negative amount indicates money owed by the account holder.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
