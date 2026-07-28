---
title: ephemeral_key
page_id: schema-ephemeral-key-6e3d666f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# ephemeral_key

```yaml
{"title": "EphemeralKey", "required": ["created", "expires", "id", "livemode", "object"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expires": {"type": "integer", "description": "Time at which the key will expire. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["ephemeral_key"]}, "secret": {"maxLength": 5000, "type": "string", "description": "The key's secret. You can use this value to make authorized requests to the Stripe API."}}, "description": "", "x-expandableFields": [], "x-resourceId": "ephemeral_key"}
```
