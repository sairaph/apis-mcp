---
title: v2.billing.meter_event_session
page_id: schema-v2-billing-meter-event-session-b867e0b5
path: schemas
description: A Meter Event Session is an authentication session for the high-throughput meter event API. Meter Event Sessions provide temporary authentication tokens with expiration times, enabling secure and efficient bulk submission of usage events.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.billing.meter_event_session

A Meter Event Session is an authentication session for the high-throughput meter event API. Meter Event Sessions provide temporary authentication tokens with expiration times, enabling secure and efficient bulk submission of usage events.

```yaml
{"title": "Meter Event Session", "required": ["authentication_token", "created", "expires_at", "id", "livemode", "object"], "type": "object", "properties": {"authentication_token": {"type": "string", "description": "The authentication token for this session. Use this token when calling the\nhigh-throughput meter event API."}, "created": {"type": "string", "description": "The creation time of this session.", "format": "date-time"}, "expires_at": {"type": "string", "description": "The time at which this session expires.", "format": "date-time"}, "id": {"type": "string", "description": "The unique ID of this auth session."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value of the object field.", "enum": ["v2.billing.meter_event_session"]}}, "description": "A Meter Event Session is an authentication session for the high-throughput meter event API. Meter Event Sessions provide temporary authentication tokens with expiration times, enabling secure and efficient bulk submission of usage events."}
```
