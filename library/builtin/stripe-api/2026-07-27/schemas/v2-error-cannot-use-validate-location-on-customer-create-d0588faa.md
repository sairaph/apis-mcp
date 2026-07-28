---
title: v2.error.cannot_use_validate_location_on_customer_create
page_id: schema-v2-error-cannot-use-validate-location-on-customer-create-d0588faa
path: schemas
description: Cannot set `automatic_indirect_tax.validate_location` when initially creating a customer configuration.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.cannot_use_validate_location_on_customer_create

Cannot set `automatic_indirect_tax.validate_location` when initially creating a customer configuration.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["cannot_use_validate_location_on_customer_create"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Cannot set `automatic_indirect_tax.validate_location` when initially creating a customer configuration."}
```
