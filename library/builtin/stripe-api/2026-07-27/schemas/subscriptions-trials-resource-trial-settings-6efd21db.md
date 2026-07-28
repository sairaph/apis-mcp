---
title: subscriptions_trials_resource_trial_settings
page_id: schema-subscriptions-trials-resource-trial-settings-6efd21db
path: schemas
description: Configures how this subscription behaves during the trial period.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_trials_resource_trial_settings

Configures how this subscription behaves during the trial period.

```yaml
{"title": "SubscriptionsTrialsResourceTrialSettings", "required": ["end_behavior"], "type": "object", "properties": {"end_behavior": {"$ref": "#/components/schemas/subscriptions_trials_resource_end_behavior"}}, "description": "Configures how this subscription behaves during the trial period.", "x-expandableFields": ["end_behavior"]}
```
