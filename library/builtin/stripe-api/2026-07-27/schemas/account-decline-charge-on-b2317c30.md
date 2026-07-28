---
title: account_decline_charge_on
page_id: schema-account-decline-charge-on-b2317c30
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_decline_charge_on

```yaml
{"title": "AccountDeclineChargeOn", "required": ["avs_failure", "cvc_failure"], "type": "object", "properties": {"avs_failure": {"type": "boolean", "description": "Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification."}, "cvc_failure": {"type": "boolean", "description": "Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification."}}, "description": "", "x-expandableFields": []}
```
