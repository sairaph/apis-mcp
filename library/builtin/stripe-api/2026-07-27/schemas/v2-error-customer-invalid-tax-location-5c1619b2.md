---
title: v2.error.customer_invalid_tax_location
page_id: schema-v2-error-customer-invalid-tax-location-5c1619b2
path: schemas
description: Invalid customer tax location.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.customer_invalid_tax_location

Invalid customer tax location.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "incorrect_fields", "message", "validate_location"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["customer_invalid_tax_location"]}, "incorrect_fields": {"type": "array", "description": "A list of fields to fix on the v2 account to achieve a valid customer tax location.", "items": {"required": ["field", "type"], "type": "object", "properties": {"field": {"type": "string", "description": "Incorrect field."}, "type": {"type": "string", "description": "Reason the field is incorrect."}}}}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "validate_location": {"type": "string", "description": "The value of the request's 'automatic_indirect_tax.validate_location' field - defaults to `auto`."}}, "description": "Information about the error that occurred"}}, "description": "Invalid customer tax location."}
```
