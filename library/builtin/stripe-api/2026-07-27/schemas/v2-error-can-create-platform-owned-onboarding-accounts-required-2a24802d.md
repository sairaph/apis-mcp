---
title: v2.error.can_create_platform_owned_onboarding_accounts_required
page_id: schema-v2-error-can-create-platform-owned-onboarding-accounts-required-2a24802d
path: schemas
description: Dormant accounts cannot create accounts where requirements collector is application (this is an account takeover prevention measure).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.can_create_platform_owned_onboarding_accounts_required

Dormant accounts cannot create accounts where requirements collector is application (this is an account takeover prevention measure).

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["can_create_platform_owned_onboarding_accounts_required"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Dormant accounts cannot create accounts where requirements collector is application (this is an account takeover prevention measure)."}
```
