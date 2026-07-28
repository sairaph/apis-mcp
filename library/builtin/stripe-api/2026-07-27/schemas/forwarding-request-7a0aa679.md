---
title: forwarding.request
page_id: schema-forwarding-request-7a0aa679
path: schemas
description: |-
    Instructs Stripe to make a request on your behalf using the destination URL. The destination URL
    is activated by Stripe at the time of onboarding. Stripe verifies requests with your credentials
    provided during onboarding, and injects card details from the payment_method into the request.

    Stripe redacts all sensitive fields and headers, including authentication credentials and card numbers,
    before storing the request and response data in the forwarding Request object, which are subject to a
    30-day retention period.

    You can provide a Stripe idempotency key to make sure that requests with the same key result in only one
    outbound request. The Stripe idempotency key provided should be unique and different from any idempotency
    keys provided on the underlying third-party request.

    Forwarding Requests are synchronous requests that return a response or time out according to
    Stripe’s limits.

    Related guide: [Forward card details to third-party API endpoints](https://docs.stripe.com/payments/forwarding).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# forwarding.request

Instructs Stripe to make a request on your behalf using the destination URL. The destination URL
is activated by Stripe at the time of onboarding. Stripe verifies requests with your credentials
provided during onboarding, and injects card details from the payment_method into the request.

Stripe redacts all sensitive fields and headers, including authentication credentials and card numbers,
before storing the request and response data in the forwarding Request object, which are subject to a
30-day retention period.

You can provide a Stripe idempotency key to make sure that requests with the same key result in only one
outbound request. The Stripe idempotency key provided should be unique and different from any idempotency
keys provided on the underlying third-party request.

Forwarding Requests are synchronous requests that return a response or time out according to
Stripe’s limits.

Related guide: [Forward card details to third-party API endpoints](https://docs.stripe.com/payments/forwarding).

```yaml
{"title": "ForwardingRequest", "required": ["created", "id", "livemode", "object", "payment_method", "replacements"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["forwarding.request"]}, "payment_method": {"maxLength": 5000, "type": "string", "description": "The PaymentMethod to insert into the forwarded request. Forwarding previously consumed PaymentMethods is allowed."}, "replacements": {"type": "array", "description": "The field kinds to be replaced in the forwarded request.", "items": {"type": "string", "enum": ["card_cvc", "card_expiry", "card_number", "cardholder_name", "request_signature"], "x-stripeBypassValidation": true}}, "request_context": {"description": "Context about the request from Stripe's servers to the destination endpoint.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/forwarded_request_context"}]}, "request_details": {"description": "The request that was sent to the destination endpoint. We redact any sensitive fields.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/forwarded_request_details"}]}, "response_details": {"description": "The response that the destination endpoint returned to us. We redact any sensitive fields.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/forwarded_response_details"}]}, "url": {"maxLength": 5000, "type": "string", "description": "The destination URL for the forwarded request. Must be supported by the config.", "nullable": true}}, "description": "Instructs Stripe to make a request on your behalf using the destination URL. The destination URL\nis activated by Stripe at the time of onboarding. Stripe verifies requests with your credentials\nprovided during onboarding, and injects card details from the payment_method into the request.\n\nStripe redacts all sensitive fields and headers, including authentication credentials and card numbers,\nbefore storing the request and response data in the forwarding Request object, which are subject to a\n30-day retention period.\n\nYou can provide a Stripe idempotency key to make sure that requests with the same key result in only one\noutbound request. The Stripe idempotency key provided should be unique and different from any idempotency\nkeys provided on the underlying third-party request.\n\nForwarding Requests are synchronous requests that return a response or time out according to\nStripe’s limits.\n\nRelated guide: [Forward card details to third-party API endpoints](https://docs.stripe.com/payments/forwarding).", "x-expandableFields": ["request_context", "request_details", "response_details"], "x-resourceId": "forwarding.request"}
```
