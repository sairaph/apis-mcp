---
title: v2.error.date_of_birth_age_restriction
page_id: schema-v2-error-date-of-birth-age-restriction-b4707da6
path: schemas
description: Representative date of birth does not meet the age limit.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.date_of_birth_age_restriction

Representative date of birth does not meet the age limit.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["age_limit", "code", "message", "user_message"], "type": "object", "properties": {"age_limit": {"type": "string", "description": "Representative age should be greater than or equal to the age_limit."}, "code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["date_of_birth_age_restriction"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Representative date of birth does not meet the age limit."}
```
