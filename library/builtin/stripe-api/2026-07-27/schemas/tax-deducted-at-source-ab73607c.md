---
title: tax_deducted_at_source
page_id: schema-tax-deducted-at-source-ab73607c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_deducted_at_source

```yaml
{"title": "TaxDeductedAtSource", "required": ["id", "object", "period_end", "period_start", "tax_deduction_account_number"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["tax_deducted_at_source"]}, "period_end": {"type": "integer", "description": "The end of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.", "format": "unix-time"}, "period_start": {"type": "integer", "description": "The start of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.", "format": "unix-time"}, "tax_deduction_account_number": {"maxLength": 5000, "type": "string", "description": "The TAN that was supplied to Stripe when TDS was assessed"}}, "description": "", "x-expandableFields": []}
```
