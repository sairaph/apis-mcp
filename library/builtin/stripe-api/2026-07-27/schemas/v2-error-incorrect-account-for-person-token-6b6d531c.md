---
title: v2.error.incorrect_account_for_person_token
page_id: schema-v2-error-incorrect-account-for-person-token-6b6d531c
path: schemas
description: A person token is created with one account but used on a different account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.incorrect_account_for_person_token

A person token is created with one account but used on a different account.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["incorrect_account_for_person_token"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "A person token is created with one account but used on a different account."}
```
