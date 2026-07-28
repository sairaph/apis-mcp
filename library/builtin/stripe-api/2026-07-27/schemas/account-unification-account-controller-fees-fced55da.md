---
title: account_unification_account_controller_fees
page_id: schema-account-unification-account-controller-fees-fced55da
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_unification_account_controller_fees

```yaml
{"title": "AccountUnificationAccountControllerFees", "required": ["payer"], "type": "object", "properties": {"payer": {"type": "string", "description": "A value indicating the responsible payer of a bundle of Stripe fees for pricing-control eligible products on this account. Learn more about [fee behavior on connected accounts](https://docs.stripe.com/connect/direct-charges-fee-payer-behavior).", "enum": ["account", "application", "application_custom", "application_express"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
