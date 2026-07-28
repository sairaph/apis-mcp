---
title: v2.error.mcc_invalid
page_id: schema-v2-error-mcc-invalid-afb76df0
path: schemas
description: MCC is invalid for `configuration.merchant.mcc`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.mcc_invalid

MCC is invalid for `configuration.merchant.mcc`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["mcc_invalid"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "MCC is invalid for `configuration.merchant.mcc`."}
```
