---
title: linked_account_options_common
page_id: schema-linked-account-options-common-0f40e1f3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# linked_account_options_common

```yaml
{"title": "linked_account_options_common", "type": "object", "properties": {"filters": {"$ref": "#/components/schemas/payment_flows_private_payment_methods_financial_connections_common_linked_account_options_filters"}, "permissions": {"type": "array", "description": "The list of permissions to request. The `payment_method` permission must be included.", "items": {"type": "string", "enum": ["balances", "ownership", "payment_method", "transactions"], "x-stripeBypassValidation": true}}, "prefetch": {"type": "array", "description": "Data features requested to be retrieved upon account creation.", "nullable": true, "items": {"type": "string", "enum": ["balances", "ownership", "transactions"], "x-stripeBypassValidation": true}}, "return_url": {"maxLength": 5000, "type": "string", "description": "For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app."}}, "description": "", "x-expandableFields": ["filters"]}
```
