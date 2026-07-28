---
title: v2.error.account_creation_liability_unacknowledged
page_id: schema-v2-error-account-creation-liability-unacknowledged-ca56e2c8
path: schemas
description: Account creation error - liability unacknowledged.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_creation_liability_unacknowledged

Account creation error - liability unacknowledged.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_creation_liability_unacknowledged"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Account creation error - liability unacknowledged."}
```
