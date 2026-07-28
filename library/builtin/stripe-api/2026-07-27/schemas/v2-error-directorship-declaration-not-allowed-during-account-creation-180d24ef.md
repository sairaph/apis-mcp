---
title: v2.error.directorship_declaration_not_allowed_during_account_creation
page_id: schema-v2-error-directorship-declaration-not-allowed-during-account-creation-180d24ef
path: schemas
description: Directorship declaration is not allowed during account creation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.directorship_declaration_not_allowed_during_account_creation

Directorship declaration is not allowed during account creation.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["directorship_declaration_not_allowed_during_account_creation"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Directorship declaration is not allowed during account creation."}
```
