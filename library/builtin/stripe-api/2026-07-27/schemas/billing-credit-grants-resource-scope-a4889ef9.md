---
title: billing_credit_grants_resource_scope
page_id: schema-billing-credit-grants-resource-scope-a4889ef9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_scope

```yaml
{"title": "BillingCreditGrantsResourceScope", "type": "object", "properties": {"price_type": {"type": "string", "description": "The price type that credit grants can apply to. We currently only support the `metered` price type. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `prices`.", "enum": ["metered"]}, "prices": {"type": "array", "description": "The prices that credit grants can apply to. We currently only support `metered` prices. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `price_type`.", "items": {"$ref": "#/components/schemas/billing_credit_grants_resource_applicable_price"}}}, "description": "", "x-expandableFields": ["prices"]}
```
