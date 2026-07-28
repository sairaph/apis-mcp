---
title: plan_tier
page_id: schema-plan-tier-0e960a61
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# plan_tier

```yaml
{"title": "PlanTier", "type": "object", "properties": {"flat_amount": {"type": "integer", "description": "Price for the entire tier.", "nullable": true}, "flat_amount_decimal": {"type": "string", "description": "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.", "format": "decimal", "nullable": true}, "unit_amount": {"type": "integer", "description": "Per unit price for units relevant to the tier.", "nullable": true}, "unit_amount_decimal": {"type": "string", "description": "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.", "format": "decimal", "nullable": true}, "up_to": {"type": "integer", "description": "Up to and including to this quantity will be contained in the tier.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
