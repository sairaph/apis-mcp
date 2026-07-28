---
title: subscriptions_resource_trial_settings_end_behavior
page_id: schema-subscriptions-resource-trial-settings-end-behavior-d2f14ff3
path: schemas
description: Defines how a subscription behaves when a trial ends.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_trial_settings_end_behavior

Defines how a subscription behaves when a trial ends.

```yaml
{"title": "SubscriptionsResourceTrialSettingsEndBehavior", "required": ["missing_payment_method"], "type": "object", "properties": {"missing_payment_method": {"type": "string", "description": "Indicates how the subscription should change when the trial ends if the user did not provide a payment method.", "enum": ["cancel", "create_invoice", "pause"]}}, "description": "Defines how a subscription behaves when a trial ends.", "x-expandableFields": []}
```
