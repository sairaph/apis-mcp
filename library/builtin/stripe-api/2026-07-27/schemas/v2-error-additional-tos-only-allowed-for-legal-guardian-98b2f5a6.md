---
title: v2.error.additional_tos_only_allowed_for_legal_guardian
page_id: schema-v2-error-additional-tos-only-allowed-for-legal-guardian-98b2f5a6
path: schemas
description: Additional terms of service are signed by someone other than the legal guardian.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.additional_tos_only_allowed_for_legal_guardian

Additional terms of service are signed by someone other than the legal guardian.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["additional_tos_only_allowed_for_legal_guardian"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Additional terms of service are signed by someone other than the legal guardian."}
```
