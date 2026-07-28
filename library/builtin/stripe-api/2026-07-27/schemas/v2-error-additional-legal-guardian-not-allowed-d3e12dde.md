---
title: v2.error.additional_legal_guardian_not_allowed
page_id: schema-v2-error-additional-legal-guardian-not-allowed-d3e12dde
path: schemas
description: More than one legal guardian is added to an account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.additional_legal_guardian_not_allowed

More than one legal guardian is added to an account.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["additional_legal_guardian_not_allowed"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "More than one legal guardian is added to an account."}
```
