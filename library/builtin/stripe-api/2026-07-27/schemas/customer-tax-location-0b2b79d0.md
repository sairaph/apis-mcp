---
title: customer_tax_location
page_id: schema-customer-tax-location-0b2b79d0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_tax_location

```yaml
{"title": "CustomerTaxLocation", "required": ["country", "source"], "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string", "description": "The identified tax country of the customer."}, "source": {"type": "string", "description": "The data source used to infer the customer's location.", "enum": ["billing_address", "ip_address", "payment_method", "shipping_destination"]}, "state": {"maxLength": 5000, "type": "string", "description": "The identified tax state, county, province, or region of the customer.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
