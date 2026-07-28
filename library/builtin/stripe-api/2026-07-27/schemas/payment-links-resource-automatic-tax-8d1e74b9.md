---
title: payment_links_resource_automatic_tax
page_id: schema-payment-links-resource-automatic-tax-8d1e74b9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_automatic_tax

```yaml
{"title": "PaymentLinksResourceAutomaticTax", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "If `true`, tax will be calculated automatically using the customer's location."}, "liability": {"description": "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}}, "description": "", "x-expandableFields": ["liability"]}
```
