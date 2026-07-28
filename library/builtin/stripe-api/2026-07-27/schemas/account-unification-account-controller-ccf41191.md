---
title: account_unification_account_controller
page_id: schema-account-unification-account-controller-ccf41191
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_unification_account_controller

```yaml
{"title": "AccountUnificationAccountController", "required": ["type"], "type": "object", "properties": {"fees": {"$ref": "#/components/schemas/account_unification_account_controller_fees"}, "is_controller": {"type": "boolean", "description": "`true` if the Connect application retrieving the resource controls the account and can therefore exercise [platform controls](https://docs.stripe.com/connect/platform-controls-for-standard-accounts). Otherwise, this field is null."}, "losses": {"$ref": "#/components/schemas/account_unification_account_controller_losses"}, "requirement_collection": {"type": "string", "description": "A value indicating responsibility for collecting requirements on this account. Only returned when the Connect application retrieving the resource controls the account.", "enum": ["application", "stripe"]}, "stripe_dashboard": {"$ref": "#/components/schemas/account_unification_account_controller_stripe_dashboard"}, "type": {"type": "string", "description": "The controller type. Can be `application`, if a Connect application controls the account, or `account`, if the account controls itself.", "enum": ["account", "application"]}}, "description": "", "x-expandableFields": ["fees", "losses", "stripe_dashboard"]}
```
