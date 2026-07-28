---
title: billing_credit_grants_resource_amount
page_id: schema-billing-credit-grants-resource-amount-2f0510b5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_amount

```yaml
{"title": "BillingCreditGrantsResourceAmount", "required": ["type"], "type": "object", "properties": {"monetary": {"description": "The monetary amount.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_credit_grants_resource_monetary_amount"}]}, "type": {"type": "string", "description": "The type of this amount. We currently only support `monetary` billing credits.", "enum": ["monetary"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["monetary"]}
```
