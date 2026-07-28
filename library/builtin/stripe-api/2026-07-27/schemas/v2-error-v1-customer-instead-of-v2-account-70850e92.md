---
title: v2.error.v1_customer_instead_of_v2_account
page_id: schema-v2-error-v1-customer-instead-of-v2-account-70850e92
path: schemas
description: V1 Customer ID cannot be used in V2 Account APIs.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.v1_customer_instead_of_v2_account

V1 Customer ID cannot be used in V2 Account APIs.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["v1_customer_instead_of_v2_account"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "V1 Customer ID cannot be used in V2 Account APIs."}
```
