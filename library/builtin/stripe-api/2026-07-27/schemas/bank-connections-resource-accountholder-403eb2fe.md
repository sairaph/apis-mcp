---
title: bank_connections_resource_accountholder
page_id: schema-bank-connections-resource-accountholder-403eb2fe
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_accountholder

```yaml
{"title": "BankConnectionsResourceAccountholder", "required": ["type"], "type": "object", "properties": {"account": {"description": "The ID of the Stripe account that this account belongs to. Only available when `account_holder.type` is `account`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "customer": {"description": "The ID for an Account representing a customer that this account belongs to. Only available when `account_holder.type` is `customer`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}]}}, "customer_account": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "description": "Type of account holder that this account belongs to.", "enum": ["account", "customer"]}}, "description": "", "x-expandableFields": ["account", "customer"]}
```
