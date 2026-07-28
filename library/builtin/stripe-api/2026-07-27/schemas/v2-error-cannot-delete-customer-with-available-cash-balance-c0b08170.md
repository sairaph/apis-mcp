---
title: v2.error.cannot_delete_customer_with_available_cash_balance
page_id: schema-v2-error-cannot-delete-customer-with-available-cash-balance-c0b08170
path: schemas
description: Account with Customer configuration cannot be closed because the customer has a cash balance.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.cannot_delete_customer_with_available_cash_balance

Account with Customer configuration cannot be closed because the customer has a cash balance.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["cannot_delete_customer_with_available_cash_balance"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Account with Customer configuration cannot be closed because the customer has a cash balance."}
```
