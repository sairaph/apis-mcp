---
title: customer_balance_customer_balance_settings
page_id: schema-customer-balance-customer-balance-settings-b3d9178c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_balance_customer_balance_settings

```yaml
{"title": "CustomerBalanceCustomerBalanceSettings", "required": ["reconciliation_mode", "using_merchant_default"], "type": "object", "properties": {"reconciliation_mode": {"type": "string", "description": "The configuration for how funds that land in the customer cash balance are reconciled.", "enum": ["automatic", "manual"]}, "using_merchant_default": {"type": "boolean", "description": "A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance"}}, "description": "", "x-expandableFields": []}
```
