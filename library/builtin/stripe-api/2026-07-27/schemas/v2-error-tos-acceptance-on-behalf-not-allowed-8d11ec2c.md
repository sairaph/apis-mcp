---
title: v2.error.tos_acceptance_on_behalf_not_allowed
page_id: schema-v2-error-tos-acceptance-on-behalf-not-allowed-8d11ec2c
path: schemas
description: TOS cannot be accepted on behalf of accounts when requirement collection is `stripe`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.tos_acceptance_on_behalf_not_allowed

TOS cannot be accepted on behalf of accounts when requirement collection is `stripe`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["tos_acceptance_on_behalf_not_allowed"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "TOS cannot be accepted on behalf of accounts when requirement collection is `stripe`."}
```
