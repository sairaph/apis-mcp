---
title: treasury_financial_accounts_resource_inbound_ach_toggle_settings
page_id: schema-treasury-financial-accounts-resource-inbound-ach-toggle-settings-52d40501
path: schemas
description: Toggle settings for enabling/disabling an inbound ACH specific feature
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_inbound_ach_toggle_settings

Toggle settings for enabling/disabling an inbound ACH specific feature

```yaml
{"title": "TreasuryFinancialAccountsResourceInboundAchToggleSettings", "required": ["requested", "status", "status_details"], "type": "object", "properties": {"requested": {"type": "boolean", "description": "Whether the FinancialAccount should have the Feature."}, "status": {"type": "string", "description": "Whether the Feature is operational.", "enum": ["active", "pending", "restricted"]}, "status_details": {"type": "array", "description": "Additional details; includes at least one entry when the status is not `active`.", "items": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggles_setting_status_details"}}}, "description": "Toggle settings for enabling/disabling an inbound ACH specific feature", "x-expandableFields": ["status_details"]}
```
