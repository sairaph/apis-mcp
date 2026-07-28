---
title: treasury_financial_accounts_resource_aba_toggle_settings
page_id: schema-treasury-financial-accounts-resource-aba-toggle-settings-c5efc4b2
path: schemas
description: Toggle settings for enabling/disabling the ABA address feature
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_aba_toggle_settings

Toggle settings for enabling/disabling the ABA address feature

```yaml
{"title": "TreasuryFinancialAccountsResourceAbaToggleSettings", "required": ["requested", "status", "status_details"], "type": "object", "properties": {"requested": {"type": "boolean", "description": "Whether the FinancialAccount should have the Feature."}, "status": {"type": "string", "description": "Whether the Feature is operational.", "enum": ["active", "pending", "restricted"]}, "status_details": {"type": "array", "description": "Additional details; includes at least one entry when the status is not `active`.", "items": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggles_setting_status_details"}}}, "description": "Toggle settings for enabling/disabling the ABA address feature", "x-expandableFields": ["status_details"]}
```
