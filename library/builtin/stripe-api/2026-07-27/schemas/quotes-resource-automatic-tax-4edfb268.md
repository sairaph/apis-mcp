---
title: quotes_resource_automatic_tax
page_id: schema-quotes-resource-automatic-tax-4edfb268
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_automatic_tax

```yaml
{"title": "QuotesResourceAutomaticTax", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Automatically calculate taxes"}, "liability": {"description": "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}, "provider": {"maxLength": 5000, "type": "string", "description": "The tax provider powering automatic tax.", "nullable": true}, "status": {"type": "string", "description": "The status of the most recent automated tax calculation for this quote.", "nullable": true, "enum": ["complete", "failed", "requires_location_inputs"]}}, "description": "", "x-expandableFields": ["liability"]}
```
