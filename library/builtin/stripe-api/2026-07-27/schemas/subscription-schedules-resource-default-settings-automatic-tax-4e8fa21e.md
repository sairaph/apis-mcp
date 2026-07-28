---
title: subscription_schedules_resource_default_settings_automatic_tax
page_id: schema-subscription-schedules-resource-default-settings-automatic-tax-4e8fa21e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedules_resource_default_settings_automatic_tax

```yaml
{"title": "SubscriptionSchedulesResourceDefaultSettingsAutomaticTax", "required": ["enabled"], "type": "object", "properties": {"disabled_reason": {"type": "string", "description": "If Stripe disabled automatic tax, this enum describes why.", "nullable": true, "enum": ["requires_location_inputs"]}, "enabled": {"type": "boolean", "description": "Whether Stripe automatically computes tax on invoices created during this phase."}, "liability": {"description": "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}}, "description": "", "x-expandableFields": ["liability"]}
```
