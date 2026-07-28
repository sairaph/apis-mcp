---
title: account_bacs_debit_payments_settings
page_id: schema-account-bacs-debit-payments-settings-c9d1bb23
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_bacs_debit_payments_settings

```yaml
{"title": "AccountBacsDebitPaymentsSettings", "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string", "description": "The Bacs Direct Debit display name for this account. For payments made with Bacs Direct Debit, this name appears on the mandate as the statement descriptor. Mobile banking apps display it as the name of the business. To use custom branding, set the Bacs Direct Debit Display Name during or right after creation. Custom branding incurs an additional monthly fee for the platform. The fee appears 5 business days after requesting Bacs. If you don't set the display name before requesting Bacs capability, it's automatically set as \"Stripe\" and the account is onboarded to Stripe branding, which is free.", "nullable": true}, "service_user_number": {"maxLength": 5000, "type": "string", "description": "The Bacs Direct Debit Service user number for this account. For payments made with Bacs Direct Debit, this number is a unique identifier of the account with our banking partners.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
