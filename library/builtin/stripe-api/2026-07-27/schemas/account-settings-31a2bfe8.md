---
title: account_settings
page_id: schema-account-settings-31a2bfe8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_settings

```yaml
{"title": "AccountSettings", "required": ["branding", "card_payments", "dashboard", "payments"], "type": "object", "properties": {"bacs_debit_payments": {"$ref": "#/components/schemas/account_bacs_debit_payments_settings"}, "branding": {"$ref": "#/components/schemas/account_branding_settings"}, "card_issuing": {"$ref": "#/components/schemas/account_card_issuing_settings"}, "card_payments": {"$ref": "#/components/schemas/account_card_payments_settings"}, "dashboard": {"$ref": "#/components/schemas/account_dashboard_settings"}, "invoices": {"$ref": "#/components/schemas/account_invoices_settings"}, "payments": {"$ref": "#/components/schemas/account_payments_settings"}, "payouts": {"$ref": "#/components/schemas/account_payout_settings"}, "sepa_debit_payments": {"$ref": "#/components/schemas/account_sepa_debit_payments_settings"}, "treasury": {"$ref": "#/components/schemas/account_treasury_settings"}}, "description": "", "x-expandableFields": ["bacs_debit_payments", "branding", "card_issuing", "card_payments", "dashboard", "invoices", "payments", "payouts", "sepa_debit_payments", "treasury"]}
```
