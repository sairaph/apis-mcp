---
title: billing_details
page_id: schema-billing-details-1bdb7d4e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_details

```yaml
{"title": "billing_details", "type": "object", "properties": {"address": {"description": "Billing address.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "email": {"maxLength": 5000, "type": "string", "description": "Email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Full name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "Billing phone number (including extension).", "nullable": true}, "tax_id": {"maxLength": 5000, "type": "string", "description": "Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.", "nullable": true}}, "description": "", "x-expandableFields": ["address"]}
```
