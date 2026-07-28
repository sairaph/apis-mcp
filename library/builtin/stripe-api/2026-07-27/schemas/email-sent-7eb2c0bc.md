---
title: email_sent
page_id: schema-email-sent-7eb2c0bc
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# email_sent

```yaml
{"title": "EmailSent", "required": ["email_sent_at", "email_sent_to"], "type": "object", "properties": {"email_sent_at": {"type": "integer", "description": "The timestamp when the email was sent.", "format": "unix-time"}, "email_sent_to": {"maxLength": 5000, "type": "string", "description": "The recipient's email address."}}, "description": "", "x-expandableFields": []}
```
