---
title: subscriptions
page_id: schema-subscriptions-9cac91c1
path: schemas
description: |-
    The list of subscribed events that trigger POST requests to the configured endpoint URL.
    Learn more about [webhook events](/kb/1213/webhooks#events).
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# subscriptions

The list of subscribed events that trigger POST requests to the configured endpoint URL.
Learn more about [webhook events](/kb/1213/webhooks#events).

```yaml
type: array
items:
    type: string
    enum:
        - nodeCreated
        - nodeNeedsApproval
        - nodeApproved
        - nodeKeyExpiringInOneDay
        - nodeKeyExpired
        - nodeDeleted
        - nodeSigned
        - nodeNeedsSignature
        - policyUpdate
        - userCreated
        - userNeedsApproval
        - userSuspended
        - userRestored
        - userDeleted
        - userApproved
        - userRoleUpdated
        - subnetIPForwardingNotEnabled
        - exitNodeIPForwardingNotEnabled
example:
    - nodeCreated
    - userDeleted
description: |
    The list of subscribed events that trigger POST requests to the configured endpoint URL.
    Learn more about [webhook events](/kb/1213/webhooks#events).
```
