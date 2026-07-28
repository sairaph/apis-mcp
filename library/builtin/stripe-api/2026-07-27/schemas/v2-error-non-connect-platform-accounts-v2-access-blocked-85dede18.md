---
title: v2.error.non_connect_platform_accounts_v2_access_blocked
page_id: schema-v2-error-non-connect-platform-accounts-v2-access-blocked-85dede18
path: schemas
description: Needs to use the newer API version or onboard to Connect.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.non_connect_platform_accounts_v2_access_blocked

Needs to use the newer API version or onboard to Connect.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["non_connect_platform_accounts_v2_access_blocked"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Needs to use the newer API version or onboard to Connect."}
```
