---
title: v2.error.account_create_activation_required
page_id: schema-v2-error-account-create-activation-required-0bc670f9
path: schemas
description: Platform must be activated to create connected accounts.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_create_activation_required

Platform must be activated to create connected accounts.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_create_activation_required"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Platform must be activated to create connected accounts."}
```
