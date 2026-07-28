---
title: v2.error.test_clocks_customer_limit_reached
page_id: schema-v2-error-test-clocks-customer-limit-reached-ce0d605e
path: schemas
description: Cannot add customer to a test clock that has already reached its customer limit.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.test_clocks_customer_limit_reached

Cannot add customer to a test clock that has already reached its customer limit.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["test_clocks_customer_limit_reached"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Cannot add customer to a test clock that has already reached its customer limit."}
```
