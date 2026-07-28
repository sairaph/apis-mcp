---
title: v2.error.capability_not_available_without_other_capability_in_country
page_id: schema-v2-error-capability-not-available-without-other-capability-in-country-1e8d1bc7
path: schemas
description: Requested feature is not available without also requesting a different feature in your country.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.capability_not_available_without_other_capability_in_country

Requested feature is not available without also requesting a different feature in your country.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["capability_not_available_without_other_capability_in_country"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Requested feature is not available without also requesting a different feature in your country."}
```
