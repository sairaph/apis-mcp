---
title: apple_pay_domain
page_id: schema-apple-pay-domain-3cc43f5c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# apple_pay_domain

```yaml
{"title": "ApplePayDomain", "required": ["created", "domain_name", "id", "livemode", "object"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "domain_name": {"maxLength": 5000, "type": "string"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["apple_pay_domain"]}}, "description": "", "x-expandableFields": [], "x-resourceId": "apple_pay_domain"}
```
