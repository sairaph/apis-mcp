---
title: customer_tax
page_id: schema-customer-tax-dc97c0dc
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_tax

```yaml
{"title": "CustomerTax", "required": ["automatic_tax", "provider"], "type": "object", "properties": {"automatic_tax": {"type": "string", "description": "Surfaces if automatic tax computation is possible given the current customer location information.", "enum": ["failed", "not_collecting", "supported", "unrecognized_location"]}, "ip_address": {"maxLength": 5000, "type": "string", "description": "A recent IP address of the customer used for tax reporting and tax location inference.", "nullable": true}, "location": {"description": "The identified tax location of the customer.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/customer_tax_location"}]}, "provider": {"type": "string", "description": "The tax calculation provider used for location resolution. Defaults to `stripe` when not using a [third-party provider](/tax/third-party-apps).", "enum": ["anrok", "avalara", "sphere", "stripe"]}}, "description": "", "x-expandableFields": ["location"]}
```
