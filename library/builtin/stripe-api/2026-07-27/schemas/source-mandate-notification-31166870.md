---
title: source_mandate_notification
page_id: schema-source-mandate-notification-31166870
path: schemas
description: |-
    Source mandate notifications should be created when a notification related to
    a source mandate must be sent to the payer. They will trigger a webhook or
    deliver an email to the customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_mandate_notification

Source mandate notifications should be created when a notification related to
a source mandate must be sent to the payer. They will trigger a webhook or
deliver an email to the customer.

```yaml
{"title": "SourceMandateNotification", "required": ["created", "id", "livemode", "object", "reason", "source", "status", "type"], "type": "object", "properties": {"acss_debit": {"$ref": "#/components/schemas/source_mandate_notification_acss_debit_data"}, "amount": {"type": "integer", "description": "A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount associated with the mandate notification. The amount is expressed in the currency of the underlying source. Required if the notification type is `debit_initiated`.", "nullable": true}, "bacs_debit": {"$ref": "#/components/schemas/source_mandate_notification_bacs_debit_data"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["source_mandate_notification"]}, "reason": {"maxLength": 5000, "type": "string", "description": "The reason of the mandate notification. Valid reasons are `mandate_confirmed` or `debit_initiated`."}, "sepa_debit": {"$ref": "#/components/schemas/source_mandate_notification_sepa_debit_data"}, "source": {"$ref": "#/components/schemas/source"}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the mandate notification. Valid statuses are `pending` or `submitted`."}, "type": {"maxLength": 5000, "type": "string", "description": "The type of source this mandate notification is attached to. Should be the source type identifier code for the payment method, such as `three_d_secure`."}}, "description": "Source mandate notifications should be created when a notification related to\na source mandate must be sent to the payer. They will trigger a webhook or\ndeliver an email to the customer.", "x-expandableFields": ["acss_debit", "bacs_debit", "sepa_debit", "source"], "x-resourceId": "source_mandate_notification"}
```
