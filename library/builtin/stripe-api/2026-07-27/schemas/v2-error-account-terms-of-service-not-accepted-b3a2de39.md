---
title: v2.error.account_terms_of_service_not_accepted
page_id: schema-v2-error-account-terms-of-service-not-accepted-b3a2de39
path: schemas
description: Terms of service must be accepted before adding merchant configuration.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_terms_of_service_not_accepted

Terms of service must be accepted before adding merchant configuration.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_terms_of_service_not_accepted"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Terms of service must be accepted before adding merchant configuration."}
```
