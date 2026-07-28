---
title: subscriptions_resource_pause_collection
page_id: schema-subscriptions-resource-pause-collection-e5758e91
path: schemas
description: |-
    The Pause Collection settings determine how we will pause collection for this subscription and for how long the subscription
    should be paused.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_pause_collection

The Pause Collection settings determine how we will pause collection for this subscription and for how long the subscription
should be paused.

```yaml
{"title": "SubscriptionsResourcePauseCollection", "required": ["behavior"], "type": "object", "properties": {"behavior": {"type": "string", "description": "The payment collection behavior for this subscription while paused.", "enum": ["keep_as_draft", "mark_uncollectible", "void"], "x-stripeBypassValidation": true}, "resumes_at": {"type": "integer", "description": "The time after which the subscription will resume collecting payments.", "format": "unix-time", "nullable": true}}, "description": "The Pause Collection settings determine how we will pause collection for this subscription and for how long the subscription\nshould be paused.", "x-expandableFields": []}
```
