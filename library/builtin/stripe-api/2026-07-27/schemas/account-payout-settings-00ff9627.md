---
title: account_payout_settings
page_id: schema-account-payout-settings-00ff9627
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_payout_settings

```yaml
{"title": "AccountPayoutSettings", "required": ["debit_negative_balances", "schedule"], "type": "object", "properties": {"debit_negative_balances": {"type": "boolean", "description": "A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See [Understanding Connect account balances](/connect/account-balances) for details. The default value is `false` when [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts, otherwise `true`."}, "schedule": {"$ref": "#/components/schemas/transfer_schedule"}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.", "nullable": true}}, "description": "", "x-expandableFields": ["schedule"]}
```
