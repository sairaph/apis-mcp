---
title: v2.error.invalid_azure_partner_authorization
page_id: schema-v2-error-invalid-azure-partner-authorization-fa2cb242
path: schemas
description: Error returned when a user tries to create an event destination without an Azure partner authorization.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.invalid_azure_partner_authorization

Error returned when a user tries to create an event destination without an Azure partner authorization.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["invalid_azure_partner_authorization"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Error returned when a user tries to create an event destination without an Azure partner authorization."}
```
