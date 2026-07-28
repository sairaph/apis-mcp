---
title: review
page_id: schema-review-c97ace4c
path: schemas
description: |-
    Reviews can be used to supplement automated fraud detection with human expertise.

    Learn more about [Radar](/radar) and reviewing payments
    [here](https://docs.stripe.com/radar/reviews).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# review

Reviews can be used to supplement automated fraud detection with human expertise.

Learn more about [Radar](/radar) and reviewing payments
[here](https://docs.stripe.com/radar/reviews).

```yaml
{"title": "RadarReview", "required": ["created", "id", "livemode", "object", "open", "opened_reason", "reason"], "type": "object", "properties": {"billing_zip": {"maxLength": 5000, "type": "string", "description": "The ZIP or postal code of the card used, if applicable.", "nullable": true}, "charge": {"description": "The charge associated with this review.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "closed_reason": {"type": "string", "description": "The reason the review was closed, or null if it has not yet been closed. One of `approved`, `refunded`, `refunded_as_fraud`, `disputed`, `redacted`, `canceled`, `payment_never_settled`, or `acknowledged`.", "nullable": true, "enum": ["acknowledged", "approved", "canceled", "disputed", "payment_never_settled", "redacted", "refunded", "refunded_as_fraud"]}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "ip_address": {"maxLength": 5000, "type": "string", "description": "The IP address where the payment originated.", "nullable": true}, "ip_address_location": {"description": "Information related to the location of the payment. Note that this information is an approximation and attempts to locate the nearest population center - it should not be used to determine a specific address.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/radar_review_resource_location"}]}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["review"]}, "open": {"type": "boolean", "description": "If `true`, the review needs action."}, "opened_reason": {"type": "string", "description": "The reason the review was opened. One of `rule` or `manual`.", "enum": ["manual", "rule"]}, "payment_intent": {"description": "The PaymentIntent ID associated with this review, if one exists.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}, "reason": {"maxLength": 5000, "type": "string", "description": "The reason the review is currently open or closed. One of `rule`, `manual`, `approved`, `refunded`, `refunded_as_fraud`, `disputed`, `redacted`, `canceled`, `payment_never_settled`, or `acknowledged`."}, "session": {"description": "Information related to the browsing session of the user who initiated the payment.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/radar_review_resource_session"}]}}, "description": "Reviews can be used to supplement automated fraud detection with human expertise.\n\nLearn more about [Radar](/radar) and reviewing payments\n[here](https://docs.stripe.com/radar/reviews).", "x-expandableFields": ["charge", "ip_address_location", "payment_intent", "session"], "x-resourceId": "review"}
```
