---
title: v2.error.legal_guardian_requires_existing_representative
page_id: schema-v2-error-legal-guardian-requires-existing-representative-e5668ed4
path: schemas
description: A legal guardian may not be added to the account without an existing representative.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.legal_guardian_requires_existing_representative

A legal guardian may not be added to the account without an existing representative.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["legal_guardian_requires_existing_representative"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "A legal guardian may not be added to the account without an existing representative."}
```
