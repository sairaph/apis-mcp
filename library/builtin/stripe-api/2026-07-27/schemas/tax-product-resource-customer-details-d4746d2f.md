---
title: tax_product_resource_customer_details
page_id: schema-tax-product-resource-customer-details-d4746d2f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_customer_details

```yaml
{"title": "TaxProductResourceCustomerDetails", "required": ["tax_ids", "taxability_override"], "type": "object", "properties": {"address": {"description": "The customer's postal address (for example, home or business location).", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/tax_product_resource_postal_address"}]}, "address_source": {"type": "string", "description": "The type of customer address provided.", "nullable": true, "enum": ["billing", "shipping"]}, "ip_address": {"maxLength": 5000, "type": "string", "description": "The customer's IP address (IPv4 or IPv6).", "nullable": true}, "tax_ids": {"type": "array", "description": "The customer's tax IDs (for example, EU VAT numbers).", "items": {"$ref": "#/components/schemas/tax_product_resource_customer_details_resource_tax_id"}}, "taxability_override": {"type": "string", "description": "The taxability override used for taxation.", "enum": ["customer_exempt", "none", "reverse_charge"]}}, "description": "", "x-expandableFields": ["address", "tax_ids"]}
```
