---
title: account_unification_account_controller_losses
page_id: schema-account-unification-account-controller-losses-9c9f0cbb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_unification_account_controller_losses

```yaml
{"title": "AccountUnificationAccountControllerLosses", "required": ["payments"], "type": "object", "properties": {"payments": {"type": "string", "description": "A value indicating who is liable when this account can't pay back negative balances from payments.", "enum": ["application", "stripe"]}}, "description": "", "x-expandableFields": []}
```
