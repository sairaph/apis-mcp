---
title: v2.error.duplicate_person_not_allowed
page_id: schema-v2-error-duplicate-person-not-allowed-cb635a9a
path: schemas
description: Duplicate person is added to an account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.duplicate_person_not_allowed

Duplicate person is added to an account.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "duplicate_person_id", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["duplicate_person_not_allowed"]}, "duplicate_person_id": {"type": "string", "description": "The ID of the other person with identical fields."}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Duplicate person is added to an account."}
```
