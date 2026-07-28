---
title: v2.error.test_clock_disallowed_on_live_mode
page_id: schema-v2-error-test-clock-disallowed-on-live-mode-4252d5ab
path: schemas
description: Cannot set a test clock on a livemode customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.test_clock_disallowed_on_live_mode

Cannot set a test clock on a livemode customer.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["test_clock_disallowed_on_live_mode"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Cannot set a test clock on a livemode customer."}
```
