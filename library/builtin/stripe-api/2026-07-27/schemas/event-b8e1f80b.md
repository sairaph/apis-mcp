---
title: event
page_id: schema-event-b8e1f80b
path: schemas
description: |-
    Snapshot events allow you to track and react to activity in your Stripe integration. When
    the state of another API resource changes, Stripe creates an `Event` object that contains
    all the relevant information associated with that action, including the affected API
    resource. For example, a successful payment triggers a `charge.succeeded` event, which
    contains the `Charge` in the event's data property. Some actions trigger multiple events.
    For example, if you create a new subscription for a customer, it triggers both a
    `customer.subscription.created` event and a `charge.succeeded` event.

    Configure an event destination in your account to listen for events that represent actions
    your integration needs to respond to. Additionally, you can retrieve an individual event or
    a list of events from the API.

    [Connect](https://docs.stripe.com/connect) platforms can also receive event notifications
    that occur in their connected accounts. These events include an account attribute that
    identifies the relevant connected account.

    You can access events through the [Retrieve Event API](https://docs.stripe.com/api/events#retrieve_event)
    for 30 days.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# event

Snapshot events allow you to track and react to activity in your Stripe integration. When
the state of another API resource changes, Stripe creates an `Event` object that contains
all the relevant information associated with that action, including the affected API
resource. For example, a successful payment triggers a `charge.succeeded` event, which
contains the `Charge` in the event's data property. Some actions trigger multiple events.
For example, if you create a new subscription for a customer, it triggers both a
`customer.subscription.created` event and a `charge.succeeded` event.

Configure an event destination in your account to listen for events that represent actions
your integration needs to respond to. Additionally, you can retrieve an individual event or
a list of events from the API.

[Connect](https://docs.stripe.com/connect) platforms can also receive event notifications
that occur in their connected accounts. These events include an account attribute that
identifies the relevant connected account.

You can access events through the [Retrieve Event API](https://docs.stripe.com/api/events#retrieve_event)
for 30 days.

```yaml
{"title": "NotificationEvent", "required": ["created", "data", "id", "livemode", "object", "pending_webhooks", "type"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string", "description": "The connected account that originates the event."}, "api_version": {"maxLength": 5000, "type": "string", "description": "The Stripe API version used to render `data` when the event was created. The contents of `data` never change, so this value remains static regardless of the API version currently in use. This property is populated only for events created on or after October 31, 2014.", "nullable": true}, "context": {"maxLength": 5000, "type": "string", "description": "Authentication context needed to fetch the event or related object."}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "data": {"$ref": "#/components/schemas/notification_event_data"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["event"]}, "pending_webhooks": {"type": "integer", "description": "Number of webhooks that haven't been successfully delivered (for example, to return a 20x response) to the URLs you specify."}, "request": {"description": "Information on the API request that triggers the event.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/notification_event_request"}]}, "type": {"maxLength": 5000, "type": "string", "description": "Description of the event (for example, `invoice.created` or `charge.refunded`)."}}, "description": "Snapshot events allow you to track and react to activity in your Stripe integration. When\nthe state of another API resource changes, Stripe creates an `Event` object that contains\nall the relevant information associated with that action, including the affected API\nresource. For example, a successful payment triggers a `charge.succeeded` event, which\ncontains the `Charge` in the event's data property. Some actions trigger multiple events.\nFor example, if you create a new subscription for a customer, it triggers both a\n`customer.subscription.created` event and a `charge.succeeded` event.\n\nConfigure an event destination in your account to listen for events that represent actions\nyour integration needs to respond to. Additionally, you can retrieve an individual event or\na list of events from the API.\n\n[Connect](https://docs.stripe.com/connect) platforms can also receive event notifications\nthat occur in their connected accounts. These events include an account attribute that\nidentifies the relevant connected account.\n\nYou can access events through the [Retrieve Event API](https://docs.stripe.com/api/events#retrieve_event)\nfor 30 days.", "x-expandableFields": ["data", "request"], "x-resourceId": "event"}
```
