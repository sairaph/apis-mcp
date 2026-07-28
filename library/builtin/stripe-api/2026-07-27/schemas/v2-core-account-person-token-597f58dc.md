---
title: v2.core.account_person_token
page_id: schema-v2-core-account-person-token-597f58dc
path: schemas
description: Person Tokens are single-use tokens which tokenize person information, and are used for creating or updating a Person.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.core.account_person_token

Person Tokens are single-use tokens which tokenize person information, and are used for creating or updating a Person.

```yaml
{"title": "Person Token", "required": ["created", "expires_at", "id", "livemode", "object", "used"], "type": "object", "properties": {"created": {"type": "string", "description": "Time at which the token was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.", "format": "date-time"}, "expires_at": {"type": "string", "description": "Time at which the token will expire.", "format": "date-time"}, "id": {"type": "string", "description": "Unique identifier for the token."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the token exists in live mode or the value `false` if the object exists in test mode."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value of the object field.", "enum": ["v2.core.account_person_token"]}, "used": {"type": "boolean", "description": "Determines if the token has already been used (tokens can only be used once)."}}, "description": "Person Tokens are single-use tokens which tokenize person information, and are used for creating or updating a Person."}
```
