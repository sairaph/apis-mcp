---
title: v2.error.invalid_konbini_payments_support_phone_number
page_id: schema-v2-error-invalid-konbini-payments-support-phone-number-28bc9c68
path: schemas
description: Konbini Payments Support Phone Number is Invalid.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_konbini_payments_support_phone_number

Konbini Payments Support Phone Number is Invalid.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_konbini_payments_support_phone_number"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Konbini Payments Support Phone Number is Invalid."}
```
