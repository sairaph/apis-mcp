---
title: webhook_endpoint
page_id: schema-webhook-endpoint-9d88bb58
path: schemas
description: |-
    You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be
    notified about events that happen in your Stripe account or connected
    accounts.

    Most users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.

    Related guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# webhook_endpoint

You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be
notified about events that happen in your Stripe account or connected
accounts.

Most users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.

Related guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)

```yaml
{"title": "NotificationWebhookEndpoint", "required": ["created", "enabled_events", "id", "livemode", "metadata", "object", "status", "url"], "type": "object", "properties": {"api_version": {"maxLength": 5000, "type": "string", "description": "The API version events are rendered as for this webhook endpoint.", "nullable": true}, "application": {"maxLength": 5000, "type": "string", "description": "The ID of the associated Connect application.", "nullable": true}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "description": {"maxLength": 5000, "type": "string", "description": "An optional description of what the webhook is used for.", "nullable": true}, "enabled_events": {"type": "array", "description": "The list of events to enable for this endpoint. `['*']` indicates that all events are enabled, except those that require explicit selection.", "items": {"maxLength": 5000, "type": "string"}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["webhook_endpoint"]}, "secret": {"maxLength": 5000, "type": "string", "description": "The endpoint's secret, used to generate [webhook signatures](https://docs.stripe.com/webhooks/signatures). Only returned at creation."}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the webhook. It can be `enabled` or `disabled`."}, "url": {"maxLength": 5000, "type": "string", "description": "The URL of the webhook endpoint."}}, "description": "You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be\nnotified about events that happen in your Stripe account or connected\naccounts.\n\nMost users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.\n\nRelated guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)", "x-expandableFields": [], "x-resourceId": "webhook_endpoint"}
```
