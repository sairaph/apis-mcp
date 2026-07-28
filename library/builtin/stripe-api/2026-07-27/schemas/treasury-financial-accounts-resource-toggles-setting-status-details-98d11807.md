---
title: treasury_financial_accounts_resource_toggles_setting_status_details
page_id: schema-treasury-financial-accounts-resource-toggles-setting-status-details-98d11807
path: schemas
description: Additional details on the FinancialAccount Features information.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_toggles_setting_status_details

Additional details on the FinancialAccount Features information.

```yaml
{"title": "TreasuryFinancialAccountsResourceTogglesSettingStatusDetails", "required": ["code"], "type": "object", "properties": {"code": {"type": "string", "description": "Represents the reason why the status is `pending` or `restricted`.", "enum": ["activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other"], "x-stripeBypassValidation": true}, "resolution": {"type": "string", "description": "Represents what the user should do, if anything, to activate the Feature.", "nullable": true, "enum": ["contact_stripe", "provide_information", "remove_restriction"], "x-stripeBypassValidation": true}, "restriction": {"type": "string", "description": "The `platform_restrictions` that are restricting this Feature.", "enum": ["inbound_flows", "outbound_flows"], "x-stripeBypassValidation": true}}, "description": "Additional details on the FinancialAccount Features information.", "x-expandableFields": []}
```
