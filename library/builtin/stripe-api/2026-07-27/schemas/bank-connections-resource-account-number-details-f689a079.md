---
title: bank_connections_resource_account_number_details
page_id: schema-bank-connections-resource-account-number-details-f689a079
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_account_number_details

```yaml
{"title": "BankConnectionsResourceAccountNumberDetails", "required": ["identifier_type", "status", "supported_networks"], "type": "object", "properties": {"expected_expiry_date": {"type": "integer", "description": "When the account number is expected to expire, if applicable.", "format": "unix-time", "nullable": true}, "identifier_type": {"type": "string", "description": "The type of account number associated with the account.", "enum": ["account_number", "tokenized_account_number"]}, "status": {"type": "string", "description": "Whether the account number is currently active and usable for transactions.", "enum": ["deactivated", "transactable"]}, "supported_networks": {"type": "array", "description": "The payment networks that the account number can be used for.", "items": {"type": "string", "enum": ["ach"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": []}
```
