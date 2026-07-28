---
title: v2.error.cross_border_connected_account_creation_not_allowed
page_id: schema-v2-error-cross-border-connected-account-creation-not-allowed-c10d3e39
path: schemas
description: Cross-border connected account creation is not allowed for this platform/account country combination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.cross_border_connected_account_creation_not_allowed

Cross-border connected account creation is not allowed for this platform/account country combination.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["cross_border_connected_account_creation_not_allowed"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Cross-border connected account creation is not allowed for this platform/account country combination."}
```
