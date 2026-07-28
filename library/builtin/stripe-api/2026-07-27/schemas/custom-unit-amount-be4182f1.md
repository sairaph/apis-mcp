---
title: custom_unit_amount
page_id: schema-custom-unit-amount-be4182f1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# custom_unit_amount

```yaml
{"title": "CustomUnitAmount", "type": "object", "properties": {"maximum": {"type": "integer", "description": "The maximum unit amount the customer can specify for this item.", "nullable": true}, "minimum": {"type": "integer", "description": "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.", "nullable": true}, "preset": {"type": "integer", "description": "The starting unit amount which can be updated by the customer.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
