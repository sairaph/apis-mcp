---
title: v2.error.test_clocks_advance_in_progress
page_id: schema-v2-error-test-clocks-advance-in-progress-67a7b6ac
path: schemas
description: Cannot modify a test clock that is currently advancing.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.test_clocks_advance_in_progress

Cannot modify a test clock that is currently advancing.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["test_clocks_advance_in_progress"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Cannot modify a test clock that is currently advancing."}
```
