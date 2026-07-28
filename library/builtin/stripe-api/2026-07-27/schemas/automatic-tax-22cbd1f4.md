---
title: automatic_tax
page_id: schema-automatic-tax-22cbd1f4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# automatic_tax

```yaml
{"title": "AutomaticTax", "required": ["enabled"], "type": "object", "properties": {"disabled_reason": {"type": "string", "description": "If Stripe disabled automatic tax, this enum describes why.", "nullable": true, "enum": ["finalization_requires_location_inputs", "finalization_system_error"]}, "enabled": {"type": "boolean", "description": "Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://docs.stripe.com/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices."}, "liability": {"description": "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}, "provider": {"maxLength": 5000, "type": "string", "description": "The tax provider powering automatic tax.", "nullable": true}, "status": {"type": "string", "description": "The status of the most recent automated tax calculation for this invoice.", "nullable": true, "enum": ["complete", "failed", "requires_location_inputs"]}}, "description": "", "x-expandableFields": ["liability"]}
```
